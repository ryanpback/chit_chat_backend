package handlers

import (
	"chitChat/helpers"
	"chitChat/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UsersIndex returns all users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	users, err := models.UsersAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// UsersShow retrieves a single user based on ID
func UsersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userID, _ := strconv.Atoi(id)

	user, err := models.UserFindByID(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// UsersCreate creates a new user
func UsersCreate(w http.ResponseWriter, r *http.Request) {
	request, err := decode(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if valid, err := helpers.ValidateUser(request, helpers.GetUserCreateFields()); !valid {
		respondFailedValidation(w, err)

		return
	}

	user, err := models.UserCreate(request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// UsersEdit will edit a user (passed by ID)
func UsersEdit(w http.ResponseWriter, r *http.Request) {
	request, err := decode(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if valid, err := helpers.ValidateUser(request, helpers.GetUserEditFields()); !valid {
		respondFailedValidation(w, err)

		return
	}

	params := mux.Vars(r)
	id := params["id"]
	userID, _ := strconv.Atoi(id)

	user, err := models.UserFindByID(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())

		return
	}

	updatedUser, err := models.UserEdit(user, request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, updatedUser)
}
