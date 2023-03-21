package storage

import (
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/lucasdk3/go-oqcomer-api/services"
)

var client *redis.Client

func StartStorage() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func CreateAuth(userid uint, tokenDetails *services.TokenDetails) error {
	accessExpires := time.Unix(tokenDetails.AccessExpires, 0) //converting Unix to UTC(to Time object)
	refreshExpires := time.Unix(tokenDetails.RefreshExpires, 0)
	now := time.Now()

	errAccess := client.Set(tokenDetails.AccessUuid, strconv.Itoa(int(userid)), accessExpires.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(tokenDetails.RefreshUuid, strconv.Itoa(int(userid)), refreshExpires.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
