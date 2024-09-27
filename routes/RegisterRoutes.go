package routes

import (
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()

	(UserRoutes(db, apiRouter))
	(AccountRoutes(db, apiRouter))

	return r
}
