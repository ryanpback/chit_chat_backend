package main

import (
	"net/http"

	"chitChat/httpd/handlers"
)

type handlerInfo struct {
	route       string
	handlerFunc func(http.ResponseWriter, *http.Request)
	methods     []string
}

func openRoutes() []handlerInfo {
	var routes []handlerInfo

	routes = []handlerInfo{
		{
			route:       "/login",
			handlerFunc: handlers.Login,
			methods:     []string{"POST", "OPTIONS"},
		},
		{
			route:       "/users",
			handlerFunc: handlers.UsersIndex,
			methods:     []string{"GET", "OPTIONS"},
		},
		{
			route:       "/users/{id: [0-9]+}",
			handlerFunc: handlers.UsersShow,
			methods:     []string{"GET", "OPTIONS"},
		},
		{
			route:       "/users",
			handlerFunc: handlers.UsersCreate,
			methods:     []string{"POST", "OPTIONS"},
		},
		{
			route:       "/users/{id: [0-9]+}",
			handlerFunc: handlers.UsersEdit,
			methods:     []string{"PATCH", "OPTIONS"},
		},
	}

	return routes
}
