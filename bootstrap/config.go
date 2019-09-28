package bootstrap

import (
	"database/sql"
	"os"

	// Pull in the dirver so we can use the postgres init function
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config this is a general purpose struct where we can keep all the app
// configuration items in a singleton style object
type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string
	Logger     *logrus.Logger
	DBConn     *sql.DB
}

// InitConfig this function will run and log out all the different environment
// variables if something isn't set correctly, it'll die and log the errors
func InitConfig() (Config, error) {
	godotenv.Load("../.env")

	c := Config{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBDatabase: os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
	// now that we've tried to pull the env values, let's set defaults if any of them are empty
	if c.DBHost == "" {
		c.DBHost = "localhost"
	}
	if c.DBDatabase == "" {
		c.DBDatabase = ""
	}
	if c.DBPassword == "" {
		// default password in case no password for local db
		c.DBPassword = os.Getenv("DB_PASSWORD_DEFAULT")
	}
	if c.DBPort == "" {
		c.DBPort = ""
	}
	if c.DBUsername == "" {
		c.DBUsername = ""
	}
	if c.AppPort == "" {
		c.AppPort = ""
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
	return c, nil
}

func dsn(c Config) string {
	dsn := "user=" + c.DBUsername + " dbname=" + c.DBDatabase + " sslmode=disable"

	return dsn
}
