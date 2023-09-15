package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

const DBDriver string = "postgres"

var (
	DBSource string
)

// Before using config variables, you should run this function first to load it.
func LoadConfig(envFullPath string) error {
	envData, err := godotenv.Read(envFullPath)

	if err != nil {
		return fmt.Errorf("'%s' is not env file, please check again", envFullPath)
	}

	DBSource = envData["DB_SOURCE"]

	return nil
}
