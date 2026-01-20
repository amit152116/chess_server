package redis

import (
	"context"
	"fmt"

	"github.com/amit152116/chess_server/config"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

func ConfigureRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisHost + ":" + config.Cfg.RedisPort,
		Password: config.Cfg.RedisPassword,
		DB:       0,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	Client = rdb
	fmt.Println("redis connect success")
}
