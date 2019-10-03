package handlers

import (
	"encoding/json"
	"net/http"
)

// decode takes request and returns a map of the request body
func decode(r *http.Request) (map[string]interface{}, error) {
	var request map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	r.Body.Close()

	return request, nil
}
