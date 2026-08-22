package main

import (
	"fmt"
	"redis-go/internal/redis"
)

func main() {

	fmt.Println("Running Redis File")

	redis.ConnectRedis()

	fmt.Println("Closing Redis")

}
