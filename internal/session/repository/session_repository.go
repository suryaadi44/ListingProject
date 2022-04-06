package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/suryaadi44/ListingProject/internal/session/entity"
)

type SessionRepository struct {
	db *mongo.Database
}

func (s *SessionRepository) NewSession(ctx context.Context, data Session) error {
	_, err := s.db.Collection("session").InsertOne(ctx, data)
	if err != nil {
		log.Println("[MONGO]", err.Error())
	}
	return err
}

func (s *SessionRepository) FindSession(ctx context.Context, uid string) (Session, error) {
	var result Session

	err := s.db.Collection("session").FindOne(ctx, bson.M{"_id": uid}).Decode(&result)
	if err != nil {
		log.Println("[MONGO]", err.Error())
	}

	return result, err
}

func (s *SessionRepository) DeleteSession(ctx context.Context, uid string) error {
	_, err := s.db.Collection("session").DeleteOne(ctx, bson.M{"_id": uid})
	if err != nil {
		log.Println("[MONGO]", err.Error())
	}

	return err
}

func NewSessionRepository(db *mongo.Database) *SessionRepository {
	return &SessionRepository{db: db}
}
