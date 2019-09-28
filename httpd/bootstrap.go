package main

import (
	bs "chitChat/bootstrap"
	"chitChat/handlers"

	"chitChat/models"
)

var appConfig bs.Config

func bootstrap() {
	c, err := bs.InitConfig()
	appConfig = c

	if err != nil {
		logError(err)
		logInfo(appConfig)
		panic("Someting went wrong, check your environment variables")
	}

	logInfo("Application bootstrapped with the following settings:")
	logInfo("Port: " + appConfig.AppPort)
	logInfo("Database: " + appConfig.DBDatabase)
	logInfo("DB Username: " + appConfig.DBUsername)

	// let's now hydrate a few things in the handlers package
	handlers.DBConn = appConfig.DBConn
	handlers.Log = appConfig.Logger

	// let's now hydrate a few things in the middleware package
	// middleware.Log = appConfig.Logger

	// let's now hydrate a few things in the models package
	models.DBConn = appConfig.DBConn
	models.Log = appConfig.Logger
}
