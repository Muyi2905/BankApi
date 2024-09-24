package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("", controllers.GetUser(db)).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetUserById(db)).Methods("GET")
	r.HandleFunc("/{id}", controllers.DeleteUser(db)).Methods("DELETE")
	r.HandleFunc("", controllers.CreateUser(db)).Methods("POST")
	r.HandleFunc("/{id}", controllers.UpdateUser(db)).Methods("PUT")
	return r
}
