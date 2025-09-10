package lib

import "github.com/go-redis/redis"

var Rds *redis.Client

func InitRds() {
	Rds = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
