package handlers

import (
	"encoding/json"
	"net/http"
)

// respondWithError passes structured params to respondWithJSON by creating a payload object
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(
		w,
		statusCode,
		map[string]string{"error": message})
}

// respondWithJSON streams data to the ResponseWriter
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&payload)
}

// RespondFailedValidation writes the necessary response for failed validation
func respondFailedValidation(w http.ResponseWriter, errors []map[string]string) {
	jsonString, jsonErr := json.Marshal(errors)
	if jsonErr != nil {
		respondWithError(w, http.StatusBadRequest, jsonErr.Error())

		return
	}

	respondWithError(w, http.StatusBadRequest, string(jsonString))
}
