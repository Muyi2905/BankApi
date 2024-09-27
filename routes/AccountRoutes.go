package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func AccountRoutes(db *gorm.DB, r *mux.Router) {

	r.HandleFunc("", controllers.CreateAccount(db)).Methods("POST")
	r.HandleFunc("", controllers.GetAccounts(db)).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetAccountById(db)).Methods("GET")
	r.HandleFunc("/{id}", controllers.UpdateAccount(db)).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteAccount(db)).Methods("DELETE")

}
