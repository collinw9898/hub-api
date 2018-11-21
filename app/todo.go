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
	db, openErr := sql.Open("sqlite3", "./data.db")
	checkErr(openErr)
	defer db.Close()

	// Query to get all the tasks
	statement, selectErr := db.Prepare("SELECT * FROM tasks")
	checkErr(selectErr)
	res, execErr := statement.Query()
	checkErr(execErr)

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

// Get an individual task from the database
func getTask(w http.ResponseWriter, r *http.Request) {
	// Connect to the db
	db, openErr := sql.Open("sqlite3", "./data.db")
	checkErr(openErr)
	defer db.Close()

	// Query to get the specific task
	params := mux.Vars(r)
	statement, selectErr := db.Prepare("SELECT * FROM tasks WHERE id = ?")
	checkErr(selectErr)
	res := statement.QueryRow(params["id"])

	// Create new task from result of query
	var id int
	var text string
	res.Scan(&id, &text)

	newTask := task{
		ID:   id,
		Text: text,
	}

	db.Close()

	// Return new task as JSON
	json.NewEncoder(w).Encode(newTask)
}

// Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	// Parse the request body, create a new task
	decoder := json.NewDecoder(r.Body)
	var newTask task
	decodeErr := decoder.Decode(&newTask)
	checkErr(decodeErr)

	// Connect to the db
	db, openErr := sql.Open("sqlite3", "./data.db")
	checkErr(openErr)
	defer db.Close()

	// Insert the new task
	statement, insertErr := db.Prepare("INSERT INTO tasks(text) VALUES(?)")
	checkErr(insertErr)
	_, execErr := statement.Exec(newTask.Text)
	checkErr(execErr)

	db.Close()

	// Returns the new task
	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	// Parse the request body, create a new task
	decoder := json.NewDecoder(r.Body)
	var newTask task
	decodeErr := decoder.Decode(&newTask)
	checkErr(decodeErr)

	// Connect to the db
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	defer db.Close()

	// Update the task
	params := mux.Vars(r)
	statement, insertErr := db.Prepare("UPDATE tasks SET text = ? WHERE id = ?")
	checkErr(insertErr)
	_, execErr := statement.Exec(newTask.Text, params["id"])
	checkErr(execErr)

	db.Close()

	// Return the updated task
	json.NewEncoder(w).Encode(newTask)
}

// Delete a task from the database
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Connect to the db
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	defer db.Close()

	db.Close()
}

// Error handling for db interactions
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
