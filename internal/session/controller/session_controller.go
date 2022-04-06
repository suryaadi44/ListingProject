package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/suryaadi44/ListingProject/internal/session/dto"
	. "github.com/suryaadi44/ListingProject/internal/session/service"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type SessionController struct {
	router         *mux.Router
	sessionService *SessionService
}

func (sc *SessionController) InitializeController() {
	sc.router.HandleFunc("/api/user", sc.GetUser).Methods("GET")
}

func (sc *SessionController) GetUser(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie("session_token"); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		global.NewBaseResponse(http.StatusOK, true, "Not logged in").SendResponse(&w)
		return
	}

	session, err := sc.sessionService.FindSession(r.Context(), c.Value)
	if err != nil {
		global.NewBaseResponse(http.StatusOK, true, "Not logged in").SendResponse(&w)
		return
	}

	global.NewBaseResponse(http.StatusOK, false, NewSessionResponse(session)).SendResponse(&w)
}

func NewSessionController(router *mux.Router, sessionService *SessionService) *SessionController {
	return &SessionController{
		router:         router,
		sessionService: sessionService,
	}
}
