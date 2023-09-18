package config

import (
	"os"

	"github.com/joho/godotenv"
)

const DBDriver string = "postgres"

var (
	DBSource      string
	ServerAddress string
)

var mappingEnv = map[string]*string{
	"DB_SOURCE":      &DBSource,
	"SERVER_ADDRESS": &ServerAddress,
}

// Before using config variables, you should run this function first to load it.
func LoadConfig(filenames ...string) error {
	err := godotenv.Load(filenames...)

	for key, variable := range mappingEnv {
		*variable = os.Getenv(key)
	}

	return err
}
