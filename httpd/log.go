package main

func logInfo(message interface{}) {
	appConfig.Logger.Info(message)
}

func logError(message interface{}) {
	appConfig.Logger.Error(message)
}

func logFatal(message interface{}) {
	appConfig.Logger.Fatal(message)
}
