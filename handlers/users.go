package handlers

import (
	"chitChat/models"
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
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, users)
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
		respondWithError(w, http.StatusInternalServerError, "User ID must be a number")

		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// UsersCreate creates a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	if r.Method == http.MethodOptions {
		return
	}

	request, err := decode(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user, err := models.CreateUser(request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}
