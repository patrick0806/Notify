package main

import (
	"context"
	"sync"
	"time"

	"github.com/patrick0806/orc-notify/internal/config"
	"github.com/patrick0806/orc-notify/internal/datasource"
	"github.com/patrick0806/orc-notify/internal/usecases"
	"github.com/redis/go-redis/v9"
)

const maxGoroutines = 2

func main() {
	config.Logger()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisClient := datasource.GetRedisClient()
	consumer := redisClient.Subscribe(ctx, "notify")

	var wg sync.WaitGroup
	messagesChannel := make(chan struct{}, maxGoroutines)

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		default:
			errorCount := 0
			msg, err := consumer.ReceiveMessage(ctx)
			if err != nil {
				errorCount++
				time.Sleep(time.Second * 10 * time.Duration(errorCount))
				//log error here
				if errorCount > 5 {
					cancel()
					return
				}
			}

			wg.Add(1)
			go func(ctx *context.Context, msg *redis.Message) {
				defer wg.Done()
				messagesChannel <- struct{}{}
				usecases.ConsumeNotify(ctx, msg)
				<-messagesChannel
			}(&ctx, msg)
		}
	}
}
