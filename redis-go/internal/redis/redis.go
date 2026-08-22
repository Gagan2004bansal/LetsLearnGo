package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis#docker",
		DB:       0,
		Protocol: 2,
	})

	defer rdb.Close()

	ctx := context.Background()

	err := rdb.Set(ctx, "name", "gagan", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
	hashField := []string{
		"model", "Meteor 350",
		"brand", "Royal Enfield",
		"type", "Enduro bikes",
		"price", "244900",
	}

	err = rdb.HSet(ctx, "bike:1", hashField).Err()
	if err != nil {
		panic(err)
	}

	res, err := rdb.HGet(ctx, "bike:1", "model").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	resFull, err := rdb.HGetAll(ctx, "bike:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(resFull)
}
