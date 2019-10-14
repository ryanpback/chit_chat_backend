package main

import (
	"chitChat/httpd/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func runApp() error {
	r := mux.NewRouter()

	for _, handlerOpts := range openRoutes() {
		r.HandleFunc(handlerOpts.route, middleware.HandleCors(handlerOpts.handlerFunc)).Methods(handlerOpts.methods...)
	}

	logInfo("Starting web api at port " + appConfig.AppPort)

	logFatal(http.ListenAndServe(":"+appConfig.AppPort, r))

	return nil
}
