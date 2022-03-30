package repository

import (
	"context"
	"log"

	"github.com/suryaadi44/ListingProject/internal/homepage/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListingsRepository struct {
	db  *mongo.Database
	ctx context.Context
}

func (l ListingsRepository) ViewBriefListings(limit int64) (entity.Listings, error) {
	findOptions := options.Find()
	findOptions.SetLimit(limit)

	result, err := l.db.Collection("listings").Find(l.ctx, bson.M{}, findOptions)
	if err != nil {
		log.Println("[DB]", err.Error())
		return nil, err
	}

	var listings entity.Listings
	for result.Next(l.ctx) {
		var listing entity.Listing
		err := result.Decode(&listing)
		if err != nil {
			log.Println("[DB]", err.Error())
		}

		listings = append(listings, listing)
	}

	if err := result.Err(); err != nil {
		log.Println("[DB]", err.Error())
		return nil, err
	}
	result.Close(l.ctx)

	return listings, nil
}

func (l ListingsRepository) CountCollection(collection string) (int64, error) {
	itemCount, err := l.db.Collection(collection).CountDocuments(l.ctx, bson.M{})
	if err != nil {
		log.Println("[DB]", err.Error())
		return 0, err
	}

	return itemCount, nil
}

func NewListingRepository(db *mongo.Database) *ListingsRepository {
	return &ListingsRepository{db: db}
}
