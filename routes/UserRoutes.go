package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
)

func UserRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("", controllers.GetUser).Methods("GET")
	r.HandleFunc("/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/{id}", controllers.UpdateUser).Methods("PUT")
	return r
}
