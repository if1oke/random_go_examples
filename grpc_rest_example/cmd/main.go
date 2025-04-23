package main

import (
	"grpc_rest/internal/app"
)

func main() {
	application := app.NewApp()
	application.Run()
}
