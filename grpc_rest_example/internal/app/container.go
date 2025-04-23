package app

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpc2 "grpc_rest/internal/api/grpc"
	"grpc_rest/internal/api/rest"
	v1 "grpc_rest/pkg/api/grpc/v1"
	"log"
	"net"
	"net/http"
)

type App struct {
	grpcServer *grpc.Server
	httpServer *http.ServeMux
}

func NewApp() *App {
	grpcServer := grpc.NewServer()
	httpServer := http.NewServeMux()

	return &App{
		grpcServer: grpcServer,
		httpServer: httpServer,
	}
}

func (a *App) Run() {
	errCh := make(chan error, 2)

	// GPRC
	go (func() {
		listener, err := net.Listen("tcp", "127.0.0.1:8000")
		if err != nil {
			errCh <- fmt.Errorf("gRPC listen error: %w", err)
		}

		userHandler := grpc2.NewUserHandler()
		v1.RegisterUserServiceServer(a.grpcServer, userHandler)
		reflection.Register(a.grpcServer)

		fmt.Println("GRPC server listening on 127.0.0.1:8000")

		err = a.grpcServer.Serve(listener)
		if err != nil {
			errCh <- fmt.Errorf("gRPC serve error: %w", err)
		}
	})()

	// REST
	go (func() {
		userHandler := rest.UserHandler{}
		a.httpServer.HandleFunc("/", userHandler.Get)
		a.httpServer.HandleFunc("/create", userHandler.Create)

		fmt.Println("HTTP server listening on 127.0.0.1:8080")

		err := http.ListenAndServe("127.0.0.1:8080", a.httpServer)
		if err != nil {
			errCh <- fmt.Errorf("HTTP serve error: %w", err)
		}
	})()

	if err := <-errCh; err != nil {
		log.Fatal(err)
	}
}
