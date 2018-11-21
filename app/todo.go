package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// Allows the use of sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type tasks []task

// Gets all the tasks in the database
func getTasks(w http.ResponseWriter, r *http.Request) {
	// Connect to the db
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	defer db.Close()

	// Query to get all the tasks
	statement, err := db.Prepare("SELECT * FROM tasks")
	checkErr(err)
	res, err := statement.Query()
	checkErr(err)

	// Insert all the tasks retrieved into a slice
	var allTasks tasks
	for res.Next() {
		var id int
		var text string
		res.Scan(&id, &text)

		newTask := task{
			ID:   id,
			Text: text,
		}
		allTasks = append(allTasks, newTask)
	}

	db.Close()

	// Return the slice as JSON
	json.NewEncoder(w).Encode(allTasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	// Connect to the db
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	defer db.Close()

	// Query to get the specific task
	params := mux.Vars(r)
	statement, err := db.Prepare("SELECT * FROM tasks WHERE id=" + params["id"])
	checkErr(err)
	res, err := statement.Query()
	checkErr(err)

	// Create new task from result of query
	var id int
	var text string
	res.Next() // I should probably figure out why this is needed lol
	res.Scan(&id, &text)

	newTask := task{
		ID:   id,
		Text: text,
	}

	db.Close()

	// Return new task as JSON
	json.NewEncoder(w).Encode(newTask)
}

// Error handling for db interactions
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
