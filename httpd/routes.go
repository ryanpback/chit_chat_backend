package main

import (
	"net/http"

	"chitChat/httpd/handlers"
)

type handlerInfo struct {
	route   string
	handler func(http.ResponseWriter, *http.Request)
	methods []string
}

func openRoutes() []handlerInfo {
	var routes []handlerInfo

	routes = []handlerInfo{
		{
			route:   "/users",
			handler: handlers.UsersIndex,
			methods: []string{"GET", "OPTIONS"},
		},
		{
			route:   "/users/{id}",
			handler: handlers.UsersShow,
			methods: []string{"GET", "OPTIONS"},
		},
		{
			route:   "/users",
			handler: handlers.UsersCreate,
			methods: []string{"POST", "OPTIONS"},
		},
	}

	return routes
}
