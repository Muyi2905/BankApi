package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
)

func AccountRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/account", controllers.CreateAccount).Methods("POST")
	r.HandleFunc("/account", controllers.GetAccounts).Methods("GET")
	r.HandleFunc("/account/{id}", controllers.GetAccountById).Methods("GET")
	r.HandleFunc("/account/{id}", controllers.UpdateAccount).Methods("PUT")
	r.HandleFunc("/account/{id}", controllers.DeleteAccount).Methods("DELETE")

	return r
}
