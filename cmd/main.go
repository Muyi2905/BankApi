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

var DB *gorm.DB

func InitDb() {

	dsn := "postgresql://BACKEND_owner:BV1eyMtC5sRr@ep-proud-voice-a5uhfcga.us-east-2.aws.neon.tech/BACKEND?sslmode=require"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}

	fmt.Println("Database connection successful")
}

func main() {

	InitDb()

	if err := DB.AutoMigrate(&models.User{}, &models.Account{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	r := routes.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8080", r))
}
