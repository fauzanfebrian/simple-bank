package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/joho/godotenv"
)

const DB_DRIVER string = "postgres"

var (
	DB_SOURCE string
)

func init() {
	if err := godotenv.Load(filepath.Join(util.GetProjectPath(), ".env")); err != nil {
		fmt.Println("config error:", err)
		return
	}

	DB_SOURCE = os.Getenv("DB_SOURCE")
}
