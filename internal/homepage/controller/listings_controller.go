package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	. "github.com/suryaadi44/ListingProject/internal/homepage/service"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type ListingsController struct {
	router *mux.Router
	ls     ListingsService
}

func (lc *ListingsController) InitializeController() {
	lc.router.HandleFunc("/", lc.homepageController)
}

func (lc *ListingsController) homepageController(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("web/template/index/index.html"))

	var err = tmpl.Execute(w, nil)

	if err != nil {
		global.NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}
}

func NewController(router *mux.Router, ls ListingsService) *ListingsController {
	return &ListingsController{router: router, ls: ls}
}
