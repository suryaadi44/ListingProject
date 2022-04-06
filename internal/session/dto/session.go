package dto

import (
	. "github.com/suryaadi44/ListingProject/internal/session/entity"
)

type SessionResponse struct {
	Username string `json:"username"`
}

func NewSessionResponse(session Session) SessionResponse {
	return SessionResponse{
		Username: session.Username,
	}
}
