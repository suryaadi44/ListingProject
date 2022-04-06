package service

import (
	"context"

	. "github.com/suryaadi44/ListingProject/internal/session/entity"
	. "github.com/suryaadi44/ListingProject/internal/session/repository"
)

type SessionService struct {
	repository *SessionRepository
}

func (service *SessionService) NewSession(ctx context.Context, data Session) error {
	return service.repository.NewSession(ctx, data)
}

func (service *SessionService) FindSession(ctx context.Context, uid string) (Session, error) {
	return service.repository.FindSession(ctx, uid)
}

func (service *SessionService) DeleteSession(ctx context.Context, uid string) error {
	return service.repository.DeleteSession(ctx, uid)
}

func NewSessionService(repository *SessionRepository) *SessionService {
	return &SessionService{repository: repository}
}
