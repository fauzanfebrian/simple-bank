package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/fauzanfebrian/simplebank/config"
	"github.com/fauzanfebrian/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	envPath := filepath.Join(util.GetProjectPath(), ".env")

	config.LoadConfig(envPath)

	testDb, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
