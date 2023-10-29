package db

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/fauzanfebrian/simplebank/util"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	envPath := path.Join(util.GetProjectPath(), ".env")
	config, err := util.LoadConfig(envPath)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot load config: %s", err))
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal().Err(fmt.Errorf("Can't connect to db: %s", err))
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
