package main

import (
	"golang-twitter-clone/api"
	"golang-twitter-clone/database"
	"golang-twitter-clone/migrations"
	log "github.com/sirupsen/logrus"
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var ctx = context.Background()

func main() {

	database.InitDatabase()
	migrations.Migrate()
	api.InitializeRouter()
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Server started")

	cashErr := cache.Set(ctx, "name", "monzer", time.Second*5).Err()
	if cashErr != nil {
		panic(cashErr)
	}

	// run an infinite loop
	for {
		val, err := cache.Get(ctx, "name").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(val)
		time.Sleep(time.Second * 2)
	}

}
