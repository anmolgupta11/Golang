package redis

import "github.com/go-redis/redis"

func RedisClient() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := pingRedis(redisClient)
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

func pingRedis(client *redis.Client) (string, error) {
	result, err := client.Ping().Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
