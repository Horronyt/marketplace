package service

import (
	"github.com/Horronyt/marketplace"
	"github.com/Horronyt/marketplace/pkg/repository"
)

type Authorization interface {
	CreateUser(user marketplace.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Listing interface {
	Create(userId int, listinng marketplace.Listing) (int, error)
	GetAll(userId int) ([]marketplace.ListingOutputFormat, error)
	GetAllAnonymously() ([]marketplace.ListingOutputFormatAnon, error)
}

type Service struct {
	Authorization
	Listing
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Listing:       NewListingService(repos.Listing),
	}
}
