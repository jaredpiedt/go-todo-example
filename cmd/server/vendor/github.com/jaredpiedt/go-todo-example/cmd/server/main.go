package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jaredpiedt/go-todo-example"
	"github.com/jaredpiedt/go-todo-example/mysql"
)

var s todo.Store

func main() {
	// Connect to our database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Initialize our store
	s, err = mysql.NewStore(db)
	if err != nil {
		panic(err)
	}

	// Create our router
	r := mux.NewRouter()
	r.HandleFunc("/items", CreateItem).Methods(http.MethodPost)
	r.HandleFunc("/items/{itemID:[0-9]+}", DeleteItem).Methods(http.MethodDelete)
	r.HandleFunc("/items/{itemID:[0-9]+}", FindItemByID).Methods(http.MethodGet)
	r.HandleFunc("/items/{itemID:[0-9]+}", UpdateItem).Methods(http.MethodPut)

	// Start the server
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), r)
}
