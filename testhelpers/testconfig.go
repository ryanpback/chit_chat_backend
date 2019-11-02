package testhelpers

import (
	"database/sql"
	"os"

	// Pull in the driver so we can use the postgres init function
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var testConfig TestConfig

// TestConfig this is a general purpose struct where we can keep the
// test configuration items in a singleton style object
type TestConfig struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string
	Logger     *logrus.Logger
	DBConn     *sql.DB
}

// InitTestConfig this function will run and log out all the different environment
// variables if something isn't set correctly, it'll die and log the errors
func InitTestConfig() (TestConfig, error) {
	godotenv.Load("../.env.test")

	c := TestConfig{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}

	// initialize logging to stdout
	c.Logger = logrus.New()
	c.Logger.Out = os.Stdout

	db, err := sql.Open("postgres", dsn(c))
	if err != nil {
		return c, err
	}

	if err = db.Ping(); err != nil {
		return c, err
	}

	c.DBConn = db
	testConfig = c

	return c, nil
}

func dsn(c TestConfig) string {
	dsn := "user=" + c.DBUsername + " dbname=" + c.DBDatabase + " sslmode=disable"

	return dsn
}
