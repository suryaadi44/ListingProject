package controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type PageController struct {
	router *mux.Router
}

func (pg *PageController) InitializeController() {
	pg.router.HandleFunc("/", pg.homepageController)
}

func (pg *PageController) homepageController(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("web/template/index/index.html"))

	var err = tmpl.Execute(w, nil)

	if err != nil {
		global.NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}
}

func NewPageController(router *mux.Router) *PageController {
	return &PageController{router: router}
}
