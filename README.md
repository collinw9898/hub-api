# hub-api

The backend for a project I'm working on, called Hub.

This api serves a couple purposes:

1.) Help manage things related to school (tasks, homework) and other misc. things (bookmarks)

2.) Further my understanding of Go

3.) Further my understanding of RESTful APIs

# Setup

The only setup required is creating the database. The database being used is SQLite3, so that will need to be installed.

How to create the database:

1.) `sqlite3 data.db` (you can name it whatever you want, I like to just call it 'data')

2.) `.read scripts/createDatabase.sql` (creates the tables needed)

# Running

`go run main.go` will start the API