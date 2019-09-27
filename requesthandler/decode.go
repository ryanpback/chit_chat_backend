package requesthandler

import (
	"encoding/json"
	"net/http"
)

// Decode takes request and returns a map of the request body
func Decode(r *http.Request, value interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(value); err != nil {
		return err
	}

	r.Body.Close()

	return nil
}
