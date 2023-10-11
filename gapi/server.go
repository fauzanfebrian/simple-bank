package gapi

import (
	"fmt"

	db "github.com/fauzanfebrian/simplebank/db/sqlc"
	"github.com/fauzanfebrian/simplebank/pb"
	"github.com/fauzanfebrian/simplebank/token"
	"github.com/fauzanfebrian/simplebank/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	pb.UnimplementedSimplebankServer
	store      db.Store
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create tokenMaker: %w", err)
	}

	gin.SetMode(config.GinMode)

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
