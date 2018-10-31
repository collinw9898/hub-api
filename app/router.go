package app

import (
	"hub-api/app/bookmarks"

	"github.com/gorilla/mux"
)

// NewRouter returns the router for the api
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Bookmarks routes
	router.HandleFunc("/api/bookmarks", bookmarks.GetBookmarks).Methods("GET")
	router.HandleFunc("/api/bookmarks/{id}", bookmarks.GetBookmark).Methods("GET")
	router.HandleFunc("/api/bookmarks", bookmarks.CreateBookmark).Methods("POST")
	router.HandleFunc("/api/bookmarks/{id}", bookmarks.UpdateBookmark).Methods("PUT")
	router.HandleFunc("/api/bookmarks/{id}", bookmarks.DeleteBookmark).Methods("DELETE")

	return router
}
