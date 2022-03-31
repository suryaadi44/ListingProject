package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	. "github.com/suryaadi44/ListingProject/internal/listings/service"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type ListingsController struct {
	router *mux.Router
	ls     ListingsService
}

func (lc *ListingsController) InitializeController() {
	lc.router.HandleFunc("/api/listings", lc.fetchController)
}

func (lc *ListingsController) fetchController(w http.ResponseWriter, r *http.Request) {
	queryVar := r.URL.Query()

	limit := queryVar.Get("limit")
	if limit == "" {
		limit = "8"
	}

	page := queryVar.Get("page")
	if page == "" {
		page = "1"
	}

	limitConv, _ := strconv.ParseInt(limit, 10, 64)
	pageConv, _ := strconv.ParseInt(page, 10, 64)

	listingsList, err := lc.ls.GetBriefListings(r.Context(), limitConv, pageConv)
	if err != nil {
		global.NewBaseResponse(http.StatusInternalServerError, true, err.Error()).SendResponse(&w)
		return
	}

	global.NewBaseResponse(http.StatusOK, false, listingsList).SendResponse(&w)
}

func NewController(router *mux.Router, ls ListingsService) *ListingsController {
	return &ListingsController{router: router, ls: ls}
}
