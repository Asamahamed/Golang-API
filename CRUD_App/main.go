package main

import (
	"Crud_Api/config"
	_ "Crud_Api/docs"
	"Crud_Api/middleware"
	"Crud_Api/models"
	"Crud_Api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Connect()

	config.DB.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	routes.InitializeUserRoutes(r)

	r.Use(middleware.LoggingMiddleware)

	// Register the Swagger handler
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
