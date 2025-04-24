package main

import (
	"fmt"
	"net/http"
	transport "psql/internal/handler/http"
	"psql/internal/repository/postgres"
	"psql/internal/usecase"
)

func main() {
	dbc, err := postgres.NewClient()
	if err != nil {
		panic(err)
	}

	txManager := postgres.NewTxManager(dbc.Pool)

	accRepo := postgres.NewAccountRepo(dbc.Pool)
	accUseCase := usecase.NewAccountUseCase(accRepo, txManager)
	accHandler := transport.NewAccountHandler(accUseCase)

	roomRepo := postgres.NewRoomRepo(dbc.Pool)
	roomUseCase := usecase.NewRoomUseCase(roomRepo, txManager)
	roomHandler := transport.NewRoomHandler(roomUseCase)

	mux := transport.InitHandlers(accHandler, roomHandler)
	fmt.Println("HTTP server on :8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
