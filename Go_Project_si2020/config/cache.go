package config

import (
	"log"

	"github.com/go-redis/redis"
)

var (
	Cache *redis.Client
)

func ConnectCache() {

	Cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}

func PingCache() {

	if _, err := Cache.Ping().Result(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("cache	|	pong")
	}

}
