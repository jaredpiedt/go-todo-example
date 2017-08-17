package main

import (
	"encoding/json"
	"github.com/jaredpiedt/go-todo-example"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateItem will create a new todo item in our database.
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var i todo.Item
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Insert the item into the database
	i, err = s.CreateItem(i)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := map[string]interface{}{
		"item": i,
	}

	jsonResponse(w, http.StatusCreated, res)
}

// DeleteItem will remove a todo item from our database.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["itemID"] == "" {
		jsonResponse(w, http.StatusBadRequest, "an item id must be provided via the url")
		return
	}

	// Delete the item in the database
	err := s.DeleteItemByID(vars["itemID"])
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// FindItemByID will return one item from our database with the provided ID.
func FindItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["itemID"] == "" {
		jsonResponse(w, http.StatusBadRequest, "an item id must be provided via the url")
		return
	}

	i, err := s.FindItemByID(vars["itemID"])
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
	}

	res := map[string]interface{}{
		"item": i,
	}

	jsonResponse(w, http.StatusOK, res)
}

// UpdateItem will update an existing item in the database.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var i todo.Item
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Update the client in the databse
	err = s.UpdateItemByID(vars["itemID"], i)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := map[string]interface{}{
		"item": i,
	}

	jsonResponse(w, http.StatusOK, res)
}

// jsonResponse will marshal an interface to JSON, set the
// appropriate headers, and set the status code given.
func jsonResponse(w http.ResponseWriter, status int, e interface{}) {
	data, err := json.Marshal(e)
	if err != nil {
		w.Write([]byte("invalid json response"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
