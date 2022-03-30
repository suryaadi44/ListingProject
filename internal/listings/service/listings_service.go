package service

import (
	"context"
	"log"

	"github.com/suryaadi44/ListingProject/internal/listings/dto"
	. "github.com/suryaadi44/ListingProject/internal/listings/repository"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type ListingsService struct {
	lr ListingsRepository
}

func (ls ListingsService) GetBriefListings(ctx context.Context, limit int64, page int64) (dto.ListingsBriefResponse, error) {
	itemCount, err := ls.lr.CountCollection(ctx, "listings")
	if err != nil {
		log.Println("[Repo]", err.Error())
		return nil, err
	}

	if itemCount == 0 {
		log.Println("[Repo] Listings does not exists")
		panic(global.NewBaseResponse(494, true, "Listings does not exists"))
	}

	if page < 1 {
		page = 1
	}

	listings, err := ls.lr.ViewBriefListings(ctx, limit, page)
	if err != nil {
		log.Println("[Repo] Error fetching :", err.Error())
		return nil, err
	}

	return *dto.NewListingsBriefResponse(listings), nil
}

func NewListingService(lr ListingsRepository) *ListingsService {
	return &ListingsService{lr: lr}
}
