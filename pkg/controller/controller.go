package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	AuthController "github.com/suryaadi44/ListingProject/internal/auth/controller"
	AuthRepository "github.com/suryaadi44/ListingProject/internal/auth/repository"
	AuthService "github.com/suryaadi44/ListingProject/internal/auth/service"
	ListingsController "github.com/suryaadi44/ListingProject/internal/listings/controller"
	ListingsRepository "github.com/suryaadi44/ListingProject/internal/listings/repository"
	ListingsService "github.com/suryaadi44/ListingProject/internal/listings/service"
	AuthMiddleware "github.com/suryaadi44/ListingProject/internal/middleware"
	PageController "github.com/suryaadi44/ListingProject/internal/page/controller"
	SessionController "github.com/suryaadi44/ListingProject/internal/session/controller"
	SessionRepository "github.com/suryaadi44/ListingProject/internal/session/repository"
	SessionService "github.com/suryaadi44/ListingProject/internal/session/service"
	Middleware "github.com/suryaadi44/ListingProject/pkg/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeController(router *mux.Router, db *mongo.Database) {
	router.Use(Middleware.ErrorHandler)

	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("web/static/"))))

	sessionRepository := SessionRepository.NewSessionRepository(db)
	sessionService := SessionService.NewSessionService(sessionRepository)
	sessionController := SessionController.NewSessionController(router, sessionService)
	sessionController.InitializeController()

	authRepository := AuthRepository.NewUserRepository(db)
	authService := AuthService.NewUserService(authRepository, sessionService)

	middlewareService := AuthMiddleware.NewMiddlewareService(sessionService)
	authController := AuthController.NewController(router, authService, sessionService, middlewareService)
	authController.InitializeController()

	listingsRepository := ListingsRepository.NewListingRepository(db)
	listingsService := ListingsService.NewListingService(listingsRepository)
	listingsController := ListingsController.NewController(router, listingsService)
	listingsController.InitializeController()

	pageController := PageController.NewPageController(router)
	pageController.InitializeController()

	authPageController := PageController.NewAuthPageController(router)
	authPageController.InitializeController()
}
