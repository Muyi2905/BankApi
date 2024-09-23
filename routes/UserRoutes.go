package routes

import (
	"package"

	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
)

func UserRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	return r
}
