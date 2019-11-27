package handlers

import (
	"chitChat/helpers"
	"chitChat/models"
	"net/http"
)

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	request, err := decode(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if valid, err := helpers.ValidateUser(request, helpers.GetUserLoginFields()); !valid {
		respondFailedValidation(w, err)

		return
	}

	user, err := models.UserLogin(request)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())

		return
	}

	response := Response{
		"user": user,
	}
	respondWithJSON(w, http.StatusOK, response)
}
