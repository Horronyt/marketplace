package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/Horronyt/marketplace"
	"github.com/Horronyt/marketplace/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenTTL   = time.Hour * 12
	signingKey = "fsdfsdfjsdi%q3i4j5r328f92mx2"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user marketplace.User) (int, error) {
	user.Salt, _ = generateSalt()
	user.Password = generatePasswordHash(user.Password, user.Salt)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	salt, err := s.repo.GetUserSalt(username)
	if err != nil {
		return "", err
	}

	user, err := s.repo.GetUser(username, generatePasswordHash(password, salt))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func generatePasswordHash(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hash.Write([]byte(salt))

	return hex.EncodeToString(hash.Sum(nil))
}
