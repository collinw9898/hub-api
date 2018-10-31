package bookmarks

import (
	"fmt"
	"net/http"
)

type bookmark struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	Category string `jaon:"category"`
}

type bookmarks []bookmark

// GetBookmarks returns all the bookmarks currently stored
func GetBookmarks(w http.ResponseWriter, r *http.Request) {

	// json.NewEncoder(w).Encode(mockarray)
}

// GetBookmark gets a single bookmark by ID
func GetBookmark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented")
}

// CreateBookmark saves a new bookmark
func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented")
}

// UpdateBookmark updates a specific bookmark
func UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented")
}

// DeleteBookmark removes a bookmark
func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented")
}
