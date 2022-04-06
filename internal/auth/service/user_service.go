package service

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	. "github.com/suryaadi44/ListingProject/internal/auth/dto"
	. "github.com/suryaadi44/ListingProject/internal/auth/entity"
	. "github.com/suryaadi44/ListingProject/internal/auth/repository"
	. "github.com/suryaadi44/ListingProject/internal/session/entity"
	. "github.com/suryaadi44/ListingProject/internal/session/service"
	. "github.com/suryaadi44/ListingProject/pkg/utils"
)

type UserService struct {
	sessionService *SessionService
	repository     *UserRepository
}

func (s *UserService) RegisterUser(ctx context.Context, user Form) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	data := User{
		Username: user.Username,
		Password: hash,
	}

	err = s.repository.NewUser(ctx, data)

	return err
}

func (s *UserService) AuthenticateUser(ctx context.Context, user Form) (Session, error) {
	saved, err := s.repository.FindUser(ctx, user.Username)
	if err != nil {
		return Session{}, err
	}

	if !CheckPasswordHash(user.Password, saved.Password) {
		return Session{}, nil
	}

	log.Println("[Login Succes]", user.Username, "login approved")
	session := Session{
		SessionToken: uuid.NewString(),
		Username:     user.Username,
		Expire:       time.Now().Add(time.Duration(SESSION_EXPIRE_IN_SECOND) * time.Second),
	}

	err = s.sessionService.NewSession(ctx, session)
	if err != nil {
		return Session{}, err
	}

	return session, nil

}

func NewUserService(repository *UserRepository, sessionService *SessionService) *UserService {
	return &UserService{repository: repository, sessionService: sessionService}
}
