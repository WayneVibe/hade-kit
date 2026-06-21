package db

import (
	"context"
	"time"

	"github.com/setcreed/hade-kit/config"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Options *redis.Options
	Client  *redis.Client
}

func (r *Redis) Init(redisConf *config.Redis) {
	if r.Options == nil {
		r.Options = &redis.Options{
			Addr:           redisConf.GetAddr(),
			DB:             redisConf.GetDB(),
			Password:       redisConf.GetPassword(),
			PoolSize:       redisConf.GetPoolSize(),
			MaxIdleConns:   redisConf.GetMaxIdleConns(),
			MaxActiveConns: redisConf.GetMaxOpenConns(),
		}
	}
	rdb := redis.NewClient(r.Options)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	r.Client = rdb
}
