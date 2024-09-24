package routes

import (
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Define /api path prefix
	apiRouter := r.PathPrefix("/api").Subrouter()

	// Define exact paths for user and account routes
	apiRouter.Path("/user").Handler(UserRoutes(db)) // Note: use Path instead of PathPrefix
	apiRouter.Path("/account").Handler(AccountRoutes())

	return r
}
