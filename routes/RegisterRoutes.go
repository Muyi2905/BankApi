package routes

import (
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	
	apiRouter := r.PathPrefix("/api").Subrouter()

	
	apiRouter.Path("/user").Handler(UserRoutes(db))   
	apiRouter.Path("/account").Handler(AccountRoutes())

	return r
}
