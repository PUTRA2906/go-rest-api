package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi/config"
	"restapi/controller"
	"restapi/model"
	"restapi/repository"
	"restapi/router"
	"restapi/service"

	"github.com/rs/cors" // Import package CORS
)

func main() {
    // Koneksi DB & auto migrate
    config.ConnectDatabase()
    config.DB.AutoMigrate(&model.Book{})

    // MVC
    bookRepo := repository.NewBookRepository()
    bookService := service.NewBookService(bookRepo)
    bookController := controller.NewBookController(bookService)
    r := router.NewRouter(bookController)

    // CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Allow frontend origin
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })

    // Apply CORS middleware to your router
    handler := c.Handler(r)

    // Start the server with CORS applied
    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
