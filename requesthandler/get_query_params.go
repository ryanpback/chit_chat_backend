package requesthandler

import (
	"net/http"
	"net/url"
)

// GetQueryParams takes request and returns a list of query parameters
func GetQueryParams(r *http.Request) url.Values {
	return r.URL.Query()
}
