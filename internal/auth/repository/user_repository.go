package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/suryaadi44/ListingProject/internal/auth/entity"
)

type UserRepository struct {
	db *mongo.Database
}

func (u *UserRepository) NewUser(ctx context.Context, data User) error {
	_, err := u.db.Collection("user_db").InsertOne(ctx, data)
	if err != nil {
		log.Println("[MONGO]", err.Error())
	}
	return err
}

func (u *UserRepository) FindUser(ctx context.Context, user string) (User, error) {
	var result User

	err := u.db.Collection("user_db").FindOne(ctx, bson.M{"_id": user}).Decode(&result)
	if err != nil {
		log.Println("[MONGO]", err.Error())
	}

	return result, err
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}
}
