package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	UserRoutes := UserRoutes(db)
	AccountRoutes := AccountRoutes(db)
	apiRouter.PathPrefix("/user").Handler(UserRoutes)
	apiRouter.PathPrefix("/account").Handler(AccountRoutes)

	return r
}
