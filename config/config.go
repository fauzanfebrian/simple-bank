package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_SOURCE string
	DB_DRIVER string = "postgres"
)

func init() {
	godotenv.Load()

	DB_SOURCE = os.Getenv("DB_SOURCE")
}
