package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/suryaadi44/ListingProject/internal/listings/dto"
	. "github.com/suryaadi44/ListingProject/internal/listings/service"
	. "github.com/suryaadi44/ListingProject/internal/session/service"
	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

type ListingsController struct {
	router *mux.Router
	ls     *ListingsService
	ss     *SessionService
}

func (lc *ListingsController) InitializeController() {
	lc.router.HandleFunc("/api/listings", lc.fetchController)
	lc.router.HandleFunc("/search", lc.searchController)
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

func (lc *ListingsController) searchController(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("web/template/search/index.html"))

	loggedIn := true

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie("session_token"); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		loggedIn = false
	}

	queryVar := r.URL.Query()
	limit := queryVar.Get("limit")
	if limit == "" {
		limit = "8"
	}

	page := queryVar.Get("page")
	if page == "" {
		page = "1"
	}

	q := queryVar.Get("q")

	limitConv, _ := strconv.ParseInt(limit, 10, 64)
	pageConv, _ := strconv.ParseInt(page, 10, 64)

	listings, _ := lc.ls.GetListingsFromFullTextSearch(r.Context(), limitConv, pageConv, q)
	session, err := lc.ss.FindSession(r.Context(), c.Value)

	if err != nil {
		loggedIn = false
	}

	data := struct {
		LoggedIn bool
		Username string
		Listings dto.ListingsBriefResponse
	}{
		LoggedIn: loggedIn,
		Username: session.Username,
		Listings: listings,
	}

	tmpl.Execute(w, data)
}

func NewController(router *mux.Router, ls *ListingsService, ss *SessionService) *ListingsController {
	return &ListingsController{router: router, ls: ls, ss: ss}
}
