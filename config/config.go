package config

import (
	"github.com/joho/godotenv"
)

const DBDriver string = "postgres"

var (
	DBSource string
)

var mappingEnv = map[string]*string{
	"DB_SOURCE": &DBSource,
}

// Before using config variables, you should run this function first to load it.
func LoadConfig(filenames ...string) error {
	envData, err := godotenv.Read(filenames...)

	if err != nil {
		return err
	}

	for key, variable := range mappingEnv {
		*variable = envData[key]
	}

	return nil
}
