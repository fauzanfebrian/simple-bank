package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

const DB_DRIVER string = "postgres"

var (
	DB_SOURCE string
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if err := godotenv.Load(filepath.Join(basepath, "..", ".env")); err != nil {
		fmt.Println("config error:", err)
		return
	}

	DB_SOURCE = os.Getenv("DB_SOURCE")
}
