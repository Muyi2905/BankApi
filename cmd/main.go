package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/muyi2905/models"
	"github.com/muyi2905/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := "postgresql://BACKEND_owner:BV1eyMtC5sRr@ep-proud-voice-a5uhfcga.us-east-2.aws.neon.tech/BACKEND?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Database connection successful")
	return db
}

func main() {
	db := InitDb()

	if err := db.AutoMigrate(&models.User{}, &models.Account{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	r := routes.RegisterRoutes(db)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
