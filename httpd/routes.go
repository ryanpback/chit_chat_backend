package main

import (
	"net/http"

	"chitChat/handlers"
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
	}

	return routes
}
