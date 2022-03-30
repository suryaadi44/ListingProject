package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type AuthPageController struct {
	router *mux.Router
}

func (ag *AuthPageController) InitializeController() {
	ag.router.HandleFunc("/login", ag.loginHandler).Methods("GET")
	ag.router.HandleFunc("/signup", ag.signupHandler).Methods("GET")
}

func (ag *AuthPageController) loginHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("web/template/login/index.html"))

	var err = tmpl.Execute(w, nil)

	if err != nil {
		global.NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}
}

func (ag *AuthPageController) signupHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("web/template/signup/index.html"))

	var err = tmpl.Execute(w, nil)

	if err != nil {
		global.NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}
}

func NewAuthPageController(router *mux.Router) *AuthPageController {
	return &AuthPageController{router: router}
}
