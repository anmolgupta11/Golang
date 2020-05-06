package redis

import (
	"encoding/json"
	"time"
)

type RedisBook struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (b RedisBook) SetBookCache(key string) error {
	client, err := RedisClient()
	if err != nil {
		return err
	}
	json, err := json.Marshal(b)
	if err != nil {
		return err
	}
	client.Set(key, json, time.Second*1000)
	return nil
}

func GetBookCache(key string) (RedisBook, error) {
	client, err := RedisClient()
	if err != nil {
		return RedisBook{}, err
	}
	val, err := client.Get(key).Result()

	if err != nil {
		return RedisBook{}, err
	}
	b := RedisBook{}
	err = json.Unmarshal([]byte(val), &b)
	if err != nil {
		return RedisBook{}, err
	}

	return b, nil
}
