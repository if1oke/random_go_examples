package main

import (
	"context"
	"fmt"
	"kafka-basics/internal/app"
	"kafka-basics/internal/domain"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	signalChan := make(chan os.Signal, 1)
	messages := make(chan domain.Message)
	done := make(chan struct{})

	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		fmt.Println("Received an interrupt, shutting down...")
		cancel()
	}()

	consumer1 := app.NewConsumerApp(ctx, wg, messages, "consumer1")
	consumer2 := app.NewConsumerApp(ctx, wg, messages, "consumer2")
	logger := app.NewLoggerApp(ctx, wg, messages)

	wg.Add(3)

	go consumer1.Run()
	go consumer2.Run()
	go logger.Run()

	go func() {
		wg.Wait()
		close(messages)
		close(done)
	}()

	<-done
}
