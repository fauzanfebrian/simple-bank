package db

import (
	"database/sql"
	"log"
	"os"
	"path"
	"testing"

	"github.com/fauzanfebrian/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	envPath := path.Join(util.GetProjectPath(), ".env")
	config, err := util.LoadConfig(envPath)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
