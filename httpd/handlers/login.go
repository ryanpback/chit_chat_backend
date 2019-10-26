package handlers

import (
	"chitChat/models"
	"chitChat/services"
	"net/http"
)

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	request, err := decode(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if valid, err := services.ValidateUser(request, services.GetUserLoginFields()); !valid {
		respondFailedValidation(w, err)

		return
	}

	user, err := models.UserLogin(request)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
