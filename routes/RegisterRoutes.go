package routes

import (
	"github.com/gorilla/mux"
	"github.com/muyi2905/controllers"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Directly map the route
	r.HandleFunc("/api/user", controllers.CreateUser(db)).Methods("POST")

	// For now, comment out other routes
	// apiRouter := r.PathPrefix("/api").Subrouter()
	// userRoutes := UserRoutes(db)
	// accountRoutes := AccountRoutes()
	// apiRouter.PathPrefix("/user").Handler(userRoutes)
	// apiRouter.PathPrefix("/account").Handler(accountRoutes)

	return r
}
