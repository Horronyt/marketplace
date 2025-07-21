package service

import (
	"github.com/Horronyt/marketplace"
	"github.com/Horronyt/marketplace/pkg/repository"
)

type listingService struct {
	repo repository.Listing
}

func NewListingService(repo repository.Listing) *listingService {
	return &listingService{repo: repo}
}

func (s *listingService) Create(userId int, listing marketplace.Listing) (int, error) {
	return s.repo.Create(userId, listing)
}

func (s *listingService) GetAll(userId int) ([]marketplace.ListingOutputFormat, error) {
	return s.repo.GetAll(userId)
}

func (s *listingService) GetAllAnonymously() ([]marketplace.ListingOutputFormatAnon, error) {
	return s.repo.GetAllAnonymously()
}
