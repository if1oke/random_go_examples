package app

import (
	"context"
	"kafka-basics/internal/domain"
	"log"
	"sync"
)

type LoggerApp struct {
	ctx      context.Context
	wg       *sync.WaitGroup
	messages chan domain.Message
}

func NewLoggerApp(ctx context.Context, wg *sync.WaitGroup, messages chan domain.Message) *LoggerApp {
	return &LoggerApp{
		ctx:      ctx,
		wg:       wg,
		messages: messages,
	}
}

func (a *LoggerApp) Run() {
	defer a.wg.Done()

	for {
		select {
		case <-a.ctx.Done():
			log.Println("logger app stopped")
			return
		case msg, ok := <-a.messages:
			if !ok {
				log.Println("Logger: messages channel closed, exited")
				return
			}
			log.Printf("Message received: %+v\n", msg)
		}
	}
}
