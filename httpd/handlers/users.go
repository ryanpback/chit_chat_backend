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

	response := Response{
		"users": users,
	}
	respondWithJSON(w, http.StatusOK, response)
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

	response := Response{
		"user": user,
	}
	respondWithJSON(w, http.StatusOK, response)
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

	response := Response{
		"user": user,
	}
	respondWithJSON(w, http.StatusCreated, response)
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

	response := Response{
		"user": updatedUser,
	}
	respondWithJSON(w, http.StatusOK, response)
}

// UsersTypeahead will return users whose email is like the text passed in from the user
func UsersTypeahead(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	searchString := params["searchString"]

	if len(searchString) <= 3 {
		respondWithError(w, http.StatusBadRequest, "String not long enough to search with")
	}

	users, err := models.UsersTypeahead(searchString)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	response := Response{
		"users": users,
	}
	respondWithJSON(w, http.StatusOK, response)
}
