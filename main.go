package main

import (
	"fmt"
	"hub-api/app"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Setting the port for the server
	port := ":8000"

	// Create Router
	router := app.NewRouter()

	// Handling CORS
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Starting the server
	fmt.Println("Now listening at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
