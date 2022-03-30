package dto

import "github.com/suryaadi44/ListingProject/internal/listings/entity"

type ListingBriefResponse struct {
	Name         string        `json:"name"`
	Price        string        `json:"price"`
	Url          string        `json:"url"`
	ImagesUrl    string        `json:"images"`
	Address      string        `json:"address"`
	Summary      string        `json:"summary"`
	Rating       int64         `json:"rating"`
	NoOfReviews  int64         `json:"noofreviews"`
	RatingDetail RatingDetails `json:"rating_detail"`
}

type RatingDetails struct {
	Accuracy      int64 `json:"accuracy"`
	CheckIn       int64 `json:"check_in"`
	Cleanliness   int64 `json:"cleanliness"`
	Communication int64 `json:"communication"`
	Location      int64 `json:"location"`
	Value         int64 `json:"value"`
}

type ListingsBriefResponse []ListingBriefResponse

func NewListingBriefResponse(l entity.Listing) ListingBriefResponse {
	return ListingBriefResponse{
		Name:        l.Name,
		Price:       l.Price.String(),
		Url:         l.Url,
		ImagesUrl:   l.Images.PictureUrl,
		Address:     l.Address.Street,
		Summary:     l.Summary,
		Rating:      l.ReviewScores.Rating,
		NoOfReviews: l.NoOfReviews,
		RatingDetail: RatingDetails{
			Accuracy:      l.ReviewScores.Accuracy,
			CheckIn:       l.ReviewScores.CheckIn,
			Cleanliness:   l.ReviewScores.Cleanliness,
			Communication: l.ReviewScores.Communication,
			Location:      l.ReviewScores.Location,
			Value:         l.ReviewScores.Value,
		},
	}
}

func NewListingsBriefResponse(l entity.Listings) *ListingsBriefResponse {
	var listingsBriefResponse ListingsBriefResponse

	for _, each := range l {
		listing := NewListingBriefResponse(each)
		listingsBriefResponse = append(listingsBriefResponse, listing)
	}

	return &listingsBriefResponse
}
