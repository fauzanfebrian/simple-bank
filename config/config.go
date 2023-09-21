package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const DBDriver string = "postgres"

var (
	DBSource      string
	ServerAddress string
	Environment   string
	GinMode       string = gin.DebugMode
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
			variable:     &Environment,
			key:          "ENVIRONMENT",
			defaultValue: "development",
		},
	}

	for _, v := range mapLoadEnv {
		*v.variable = os.Getenv(v.key)
		if *v.variable == "" && v.defaultValue != "" {
			os.Setenv(v.key, v.defaultValue)
			*v.variable = v.defaultValue
		}
	}

	switch Environment {
	case "production":
		GinMode = gin.ReleaseMode
	case "test":
		GinMode = gin.TestMode
	default:
		GinMode = gin.DebugMode
	}
}
