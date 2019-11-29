package handlers

import (
	"chitChat/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ConversationsUser gets all conversations a particular user is part of
func ConversationsUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userID, _ := strconv.Atoi(id)

	conversations, err := models.ConversationsFindByUserID(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	response := Response{
		"conversations": conversations,
	}
	respondWithJSON(w, http.StatusOK, response)
}
