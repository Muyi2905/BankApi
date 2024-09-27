package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, r *mux.Router) {

	r.HandleFunc("/", controllers.CreateUser(db)).Methods("POST")
	r.HandleFunc("", controllers.GetUser(db)).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetUserById(db)).Methods("GET")
	r.HandleFunc("/{id}", controllers.UpdateUser(db)).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteUser(db)).Methods("DELETE")

}
