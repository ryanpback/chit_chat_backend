package requesthandler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// RequestBody type is a map[string]interface{} to handle 'json' type data
type RequestBody map[string]interface{}

// ExtractRequestBody takes request and returns a map of the request body
func ExtractRequestBody(r *http.Request) (RequestBody, error) {
	body := make(RequestBody)
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		return RequestBody{}, errors.New("No request data was found in the body")
	}

	if err == io.EOF {
		return RequestBody{}, err
	}

	r.Body.Close()

	return body, nil
}
