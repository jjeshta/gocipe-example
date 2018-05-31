// generated by gocipe 02fef3d117f1029d4142b6b7ae2d1ea6f313fd8a2f44e25333775a308c8afb37; DO NOT EDIT

package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	// Load database driver
	_ "github.com/lib/pq"
)

const (
	//EnvironmentProd represents production environment
	EnvironmentProd = "PROD"

	//EnvironmentDev represents development environment
	EnvironmentDev = "DEV"
)

var (
	// bootstrapped is a flag to prevent multiple bootstrapping
	bootstrapped = false

	// Env indicates in which environment (prod / dev) the application is running
	Env string
)

// Config represents application configuration loaded during bootstrap
type Config struct {
	DB       *sql.DB
	HTTPPort string
}

// Bootstrap loads environment variables and initializes the application
func Bootstrap() *Config {
	var config Config

	if bootstrapped {
		return nil
	}

	godotenv.Load()

	Env = os.Getenv("ENV")
	if Env == "" {
		Env = EnvironmentProd
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("Environment variable DSN must be defined. Example: postgres://user:pass@host/db?sslmode=disable")
	}

	var err error
	config.DB, err = sql.Open("postgres", dsn)
	if err == nil {
		log.Println("Connected to database successfully.")
	} else if Env == EnvironmentDev {
		log.Println("Database connection failed: ", err)
	} else {
		log.Fatal("Database connection failed: ", err)
	}

	err = config.DB.Ping()
	if err == nil {
		log.Println("Pinged database successfully.")
	} else if Env == EnvironmentDev {
		log.Println("Database ping failed: ", err)
	} else {
		log.Fatal("Database ping failed: ", err)
	}

	config.HTTPPort = os.Getenv("HTTP_PORT")
	if config.HTTPPort == "" {
		config.HTTPPort = "7000"
	}

	os.Clearenv() //prevent non-authorized access

	return &config
}