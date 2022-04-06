package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	. "github.com/suryaadi44/ListingProject/internal/auth/dto"
	. "github.com/suryaadi44/ListingProject/internal/auth/service"
	. "github.com/suryaadi44/ListingProject/internal/middleware"
	. "github.com/suryaadi44/ListingProject/internal/session/service"
	. "github.com/suryaadi44/ListingProject/pkg/dto"
	. "github.com/suryaadi44/ListingProject/pkg/utils"

	"github.com/gorilla/mux"
)

type UserController struct {
	router            *mux.Router
	userService       *UserService
	sessionService    *SessionService
	middlewareService *Middleware
}

func (u *UserController) InitializeController() {
	u.router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("web/static/"))))

	u.router.HandleFunc("/login", u.loginHandler).Methods("POST")
	u.router.HandleFunc("/signup", u.signupHandler).Methods("POST")
	u.router.HandleFunc("/logout", u.logoutHandler)
	u.router.HandleFunc("/blocked", u.blockedHandler)

	getContent := u.router.PathPrefix("/succes").Subrouter()
	getContent.Use(u.middlewareService.AuthMiddleware())
	getContent.HandleFunc("", u.succesHandler)
}

func (u *UserController) loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	payload := Form{}

	if err := decoder.Decode(&payload); err != nil {
		log.Println("[Decode] Error decoding JSON")
		NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}

	session, err := u.userService.AuthenticateUser(payload)
	if err != nil {
		log.Println("[Login Failed]", payload.Username, "login failed")
		NewBaseResponse(http.StatusUnauthorized, true, "Inccorect Username or Password").SendResponse(&w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   session.SessionToken,
		Expires: session.Expire,
	})

	NewBaseResponse(http.StatusSeeOther, false, "/succes").SendResponse(&w)
	return

}

func (u *UserController) signupHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	payload := Form{}

	if err := decoder.Decode(&payload); err != nil {
		log.Println("[Decode] Error decoding JSON")
		NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}

	if ok := IsEmailValid(payload.Username); !ok {
		NewBaseResponse(http.StatusBadRequest, true, "Email not valid").SendResponse(&w)
		return
	}

	err := u.userService.RegisterUser(payload)
	if err == nil {
		log.Println("[Sign Up Succes]", payload.Username, "created")
		NewBaseResponse(http.StatusSeeOther, false, "/login").SendResponse(&w)
		return
	}

	if !strings.Contains(err.Error(), "duplicate") {
		NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}

	log.Println("[Sign Up Failed]", payload.Username, "Already exist")
	NewBaseResponse(http.StatusOK, true, fmt.Sprintf("Accout with username %s already exist", payload.Username)).SendResponse(&w)
	return
}

func (u UserController) logoutHandler(w http.ResponseWriter, r *http.Request) {
	if storedCookie, _ := r.Cookie("session_token"); storedCookie != nil {
		u.sessionService.DeleteSession(storedCookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (u *UserController) succesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/succes.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}
}

func (u *UserController) blockedHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/blocked.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}
}

func NewController(router *mux.Router, userService *UserService, sessionService *SessionService, middlewareService *Middleware) *UserController {
	return &UserController{router: router, userService: userService, sessionService: sessionService, middlewareService: middlewareService}
}
