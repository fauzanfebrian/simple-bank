package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/fauzanfebrian/simplebank/api"
	db "github.com/fauzanfebrian/simplebank/db/sqlc"
	"github.com/fauzanfebrian/simplebank/gapi"
	"github.com/fauzanfebrian/simplebank/pb"
	"github.com/fauzanfebrian/simplebank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect db:", err)
	}

	store := db.NewStore(conn)

	RunGrpcServer(config, store)
	// RunGinServer(config, store)
}

func RunGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimplebankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot server gRPC:", err)
	}
}

func RunGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	fmt.Println("Starting server on: \"" + config.HTTPServerAddress + "\"")

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
