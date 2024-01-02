package main

import (
	"context"
	"log"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = initRedis()

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func PingRedis() {
	_, err := rdb.Ping(ctx).Result()
	if err == nil {
		log.Println("connected to redis successfully")
	} else {
		log.Panic("Cannot connect to redis")
	}
}

func GetValueRedis(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func SetValueRedis(key string, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

func CheckKeyRedis(key string) (bool, error) {
	n, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func DeleteKeyRedis(key string) error {
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}