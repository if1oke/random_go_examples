package app

import (
	"context"
	"fmt"
	"kafka-basics/internal/domain"
	"kafka-basics/internal/infrastructure/kafka"
	usecase2 "kafka-basics/internal/usecase"
	"log"
	"sync"
)

type ConsumerApp struct {
	ctx  context.Context
	wg   *sync.WaitGroup
	out  chan domain.Message
	name string
}

func NewConsumerApp(ctx context.Context, wg *sync.WaitGroup, out chan domain.Message, name string) *ConsumerApp {
	return &ConsumerApp{
		ctx:  ctx,
		wg:   wg,
		out:  out,
		name: name,
	}
}

func (a *ConsumerApp) Run() {
	defer a.wg.Done()

	consumer := kafka.NewConsumer(
		"192.168.100.191:9092",
		"base_g1",
		"demo-test",
	)
	defer consumer.Reader.Close()

	useCase := usecase2.NewMessageConsumerUseCase(consumer)
	for {
		select {
		case <-a.ctx.Done():
			fmt.Printf("Consumer %s stopped\n", a.name)
			return
		default:
			m, err := useCase.Read(a.ctx)
			if err != nil {
				log.Printf("Error reading message: %v\n", err)
				continue
			}
			a.out <- m
		}
	}
}
