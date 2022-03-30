package service

import (
	"log"

	"github.com/suryaadi44/ListingProject/internal/homepage/dto"
	. "github.com/suryaadi44/ListingProject/internal/homepage/repository"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type ListingsService struct {
	lr ListingsRepository
}

func (ls ListingsService) GetBriefListings(limit int64) (dto.ListingsBriefResponse, error) {
	itemCount, err := ls.lr.CountCollection("listings")
	if err != nil {
		log.Println("[Repo]", err.Error())
		return nil, err
	}

	if itemCount == 0 {
		log.Println("[Repo] Listings does not exists")
		panic(global.NewBaseResponse(494, true, "Listings does not exists"))
	}

	listings, err := ls.lr.ViewBriefListings(limit)
	if err != nil {
		log.Println("[Repo] Error fetching :", err.Error())
		return nil, err
	}

	return *dto.NewListingsBriefResponse(listings), nil
}

func NewListingService(lr ListingsRepository) *ListingsService {
	return &ListingsService{lr: lr}
}
