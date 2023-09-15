package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/fauzanfebrian/simplebank/config"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(config.DB_DRIVER, config.DB_SOURCE)

	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
