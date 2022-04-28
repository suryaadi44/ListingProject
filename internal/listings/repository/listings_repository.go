package repository

import (
	"context"
	"log"

	"github.com/suryaadi44/ListingProject/internal/listings/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListingsRepository struct {
	db *mongo.Database
}

func (l ListingsRepository) ViewBriefListings(ctx context.Context, limit int64, page int64) (entity.Listings, error) {
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(limit * (page - 1))

	result, err := l.db.Collection("listings").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Println("[DB]", err.Error())
		return nil, err
	}

	var listings entity.Listings
	for result.Next(ctx) {
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
	result.Close(ctx)

	return listings, nil
}

func (l ListingsRepository) FullTextSearchListings(ctx context.Context, limit int64, page int64, q string) (entity.Listings, error) {
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(limit * (page - 1))

	result, err := l.db.Collection("listings").Find(ctx, bson.M{"$text": bson.M{"$search": q}}, findOptions)
	if err != nil {
		log.Println("[DB]", err.Error())
		return nil, err
	}

	var listings entity.Listings
	for result.Next(ctx) {
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
	result.Close(ctx)

	return listings, nil
}

func (l ListingsRepository) CountCollection(ctx context.Context, collection string) (int64, error) {
	itemCount, err := l.db.Collection(collection).CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("[DB]", err.Error())
		return 0, err
	}

	return itemCount, nil
}

func NewListingRepository(db *mongo.Database) *ListingsRepository {
	return &ListingsRepository{db: db}
}
