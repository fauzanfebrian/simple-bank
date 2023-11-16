package db

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/fauzanfebrian/simplebank/util"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var testStore Store

func TestMain(m *testing.M) {
	var err error

	envPath := path.Join(util.GetProjectPath(), ".env")
	config, err := util.LoadConfig(envPath)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot load config: %s", err))
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("Can't connect to db: %s", err))
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
