package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const DBDriver string = "postgres"

var (
	DBSource      string
	ServerAddress string
	GinMode       string
)

// Before using config variables, you should run this function first to load it.
func LoadConfig(filenames ...string) {
	err := godotenv.Load(filenames...)

	if err != nil {
		fmt.Println("err load .env file:", err)
	}

	var mapLoadEnv = []struct {
		variable     *string
		key          string
		defaultValue string
	}{
		{
			variable:     &DBSource,
			key:          "DB_SOURCE",
			defaultValue: "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable",
		},
		{
			variable:     &ServerAddress,
			key:          "SERVER_ADDRESS",
			defaultValue: ":8080",
		},
		{
			variable: &GinMode,
			key:      "GIN_MODE",
		},
	}

	for _, v := range mapLoadEnv {
		*v.variable = os.Getenv(v.key)
		if *v.variable == "" && v.defaultValue != "" {
			*v.variable = v.defaultValue
		}
	}
}
