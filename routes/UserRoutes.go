package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Define routes for user
	r.HandleFunc("/", controllers.CreateUser(db)).Methods("POST")       // For POST /api/user
	r.HandleFunc("", controllers.GetUser(db)).Methods("GET")            // For GET /api/user
	r.HandleFunc("/{id}", controllers.GetUserById(db)).Methods("GET")   // For GET /api/user/{id}
	r.HandleFunc("/{id}", controllers.UpdateUser(db)).Methods("PUT")    // For PUT /api/user/{id}
	r.HandleFunc("/{id}", controllers.DeleteUser(db)).Methods("DELETE") // For DELETE /api/user/{id}

	return r
}
