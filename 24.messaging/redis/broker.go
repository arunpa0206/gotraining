package main

import (
	"github.com/go-redis/redis/v7"
)

// Broker is a helper for sending and receiving messages.
type Broker struct {
	client *redis.Client
}

// New is a constructor for Broker.
func New() Broker {
	return Broker{
		redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (b Broker) sendTask(key string, value string) {
	err := b.client.LPush(key, value).Err()
	if err != nil {
		panic(err)
	}
}

func (b Broker) receiveTask(key string) string {
	result, _ := b.client.RPop(key).Result()
	return result
}
