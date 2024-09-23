package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
)

func AccountRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("", controllers.CreateAccount).Methods("POST")
	r.HandleFunc("", controllers.GetAccounts).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetAccountById).Methods("GET")
	r.HandleFunc("/{id}", controllers.UpdateAccount).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteAccount).Methods("DELETE")

	return r
}
