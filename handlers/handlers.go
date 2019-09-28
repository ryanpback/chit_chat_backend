package handlers

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// DBConn this is here so that we can hydrate all the web handlers with the same
// DB connection that we are using in the main package
var DBConn *sql.DB

// Log this is here so we can share the same logger with the main package
var Log *logrus.Logger
