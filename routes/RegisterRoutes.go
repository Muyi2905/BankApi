package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	UserRoutes := UserRoutes()
	AccountRoutes := AccountRoutes()
	apiRouter.PathPrefix("/user").Handler(UserRoutes)
	apiRouter.PathPrefix("/account").Handler(AccountRoutes)

	return r
}
