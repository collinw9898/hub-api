package app

import (
	"github.com/gorilla/mux"
)

// NewRouter returns the router for the api
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Todo routes
	router.HandleFunc("/api/todo", getTasks).Methods("GET")
	router.HandleFunc("/api/todo/{id}", getTask).Methods("GET")
	// router.HandleFunc("/api/todo").Methods("POST")
	// router.HandleFunc("/api/todo/{id}").Methods("PUT")
	// router.HandleFunc("/api/todo/{id}").Methods("DELETE")

	return router
}
