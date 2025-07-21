package service

import "github.com/Horronyt/marketplace/pkg/repository"

type Authorization interface {
}

type Listing interface {
}

type Service struct {
	Authorization
	Listing
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
