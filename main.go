package main

import (
	"fmt"
	"log"

	"github.com/anvari1313/yaus/app"
	"github.com/anvari1313/yaus/config"
	"github.com/go-redis/redis/v7"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	cmd := client.Set("key", 12000, 0)
	err = cmd.Err()
	if err != nil {
		log.Fatal(err)
	}

	result := client.Get("key")
	v, err := result.Int64()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value is %d\n", v)
	res := client.Incr("key")
	err = res.Err()
	if err != nil {
		log.Fatal(err)
	}
	v, err = res.Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value is %d\n", v)
	result = client.Get("key")
	v, err = result.Int64()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Value is %d\n", v)

	c := config.ReadConfig("")
	app.CreateApp(c)
}
