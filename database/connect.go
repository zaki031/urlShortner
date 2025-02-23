package database

import (
	"log"
	"strconv"

	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)
//todo: add single ton patern
	var Client *redis.Client

	func Connect() {
		err := godotenv.Load();
		if err != nil {
			log.Fatalf("error loading env file");
		}
		//converting to int because os.Getenv returns a string
		db , _ := strconv.Atoi(os.Getenv("REDIS_DB"))
		Client = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:   db,
		})

	}
