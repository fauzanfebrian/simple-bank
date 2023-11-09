package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/fauzanfebrian/simplebank/api"
	db "github.com/fauzanfebrian/simplebank/db/sqlc"
	_ "github.com/fauzanfebrian/simplebank/doc/statik"
	"github.com/fauzanfebrian/simplebank/gapi"
	"github.com/fauzanfebrian/simplebank/mail"
	"github.com/fauzanfebrian/simplebank/pb"
	"github.com/fauzanfebrian/simplebank/util"
	"github.com/fauzanfebrian/simplebank/worker"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot load config: %s", err))
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot connect db: %s", err))
	}

	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	go RunGatewayServer(config, store, taskDistributor)
	go RunTaskProcessor(config, redisOpt, store)
	RunGrpcServer(config, store, taskDistributor)
	// RunGinServer(config, store)
}

func RunGrpcServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot create server: %s", err))
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterSimplebankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot create listener: %s", err))
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot start gRPC server: %s", err))
	}
}

func RunGatewayServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot create server: %s", err))
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimplebankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot register handler server"))
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFs, err := fs.New()
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot create statik: %s", err))
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFs))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot create listener: %s", err))
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	handler := gapi.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot start HTTP server: %s", err))
	}
}

func RunTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewSmtpEmailSender(config)

	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot start task processor: %s", err))
	}
}

func RunGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot create server: %s", err))
	}

	fmt.Println("Starting server on: \"" + config.HTTPServerAddress + "\"")

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(fmt.Errorf("cannot start server: %s", err))
	}
}
