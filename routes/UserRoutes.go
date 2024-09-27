package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, r *mux.Router) {

	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/", controllers.CreateUser(db)).Methods("POST")
	userRouter.HandleFunc("", controllers.GetUser(db)).Methods("GET")
	userRouter.HandleFunc("/{id}", controllers.GetUserById(db)).Methods("GET")
	userRouter.HandleFunc("/{id}", controllers.UpdateUser(db)).Methods("PUT")
	userRouter.HandleFunc("/{id}", controllers.DeleteUser(db)).Methods("DELETE")

}
