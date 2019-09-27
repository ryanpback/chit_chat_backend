package controllers

import (
	"chitChat/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UsersIndex returns all users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	if r.Method == http.MethodOptions {
		return
	}

	users, err := models.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(&users)
}

// UsersShow retrieves a single user based on ID
func UsersShow(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	if r.Method == http.MethodOptions {
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	userID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "User ID must be a number", http.StatusInternalServerError)

		return
	}

	user, err := models.GetUserByID(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	json.NewEncoder(w).Encode(&user)
}

// UsersCreate creates a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	//
}
