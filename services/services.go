package services

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// DBConn that we are using in the main package
var DBConn *sql.DB

// Log this is here so we can share the same logger with the main package
var Log *logrus.Logger

// Payload is how all requests and responses are wrapped
type Payload map[string]interface{}
