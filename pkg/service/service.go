package service

import (
	"github.com/Horronyt/marketplace"
	"github.com/Horronyt/marketplace/pkg/repository"
)

type Authorization interface {
	CreateUser(user marketplace.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Listing interface {
}

type Service struct {
	Authorization
	Listing
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
