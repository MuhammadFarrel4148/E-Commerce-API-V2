package caching

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func ConnectRedis() {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opt)

	_, err = rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatal("Gagal terhubung ke redis:", err)
	}

	log.Println("Sukses terhubung ke redis")
}

func Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		return "",  nil
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func Del(key string) error {
	return rdb.Del(ctx, key).Err()
}