package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"vibrox-echo/proto/logger"
)

type Server struct {
	logger.UnimplementedLoggerServer
}

func main() {

	if err := InitLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer func() {
		if err := Logger.Sync(); err != nil {
			panic(err)
		}
	}()

	lc := net.ListenConfig{}
	listen, err := lc.Listen(context.Background(), "tcp", ":9000") // #nosec G102 -- Intentionally listening on all interfaces.
	if err != nil {
		log.Fatal("Failed to listen on Port 9000: ", err)
	}

	grpcServer := grpc.NewServer()

	logger.RegisterLoggerServer(grpcServer, &Server{})

	log.Println("Logger Service is starting on port 9000...")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Failed to start logger service on port 9000: ", err)
	}
}
