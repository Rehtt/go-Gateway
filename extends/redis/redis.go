package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type Client struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

var RDB *Client

func Init() error {
	RDB = &Client{
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.addr") + ":" + viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		}),
		Ctx: context.Background(),
	}
	_, err := RDB.RedisClient.Ping(RDB.Ctx).Result()
	if err != nil {
		return fmt.Errorf("redis connection failed: %s", err.Error())
	}
	return err
}
