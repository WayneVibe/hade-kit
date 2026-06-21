package database

import (
	"github.com/setcreed/hade-kit/config"
	"github.com/setcreed/hade-kit/db"
)

var (
	RedisCli *db.Redis
)

func InitRedis(redisConf *config.Redis) {
	if redisConf == nil {
		return
	}

	r := db.Redis{}
	r.Init(redisConf)
	RedisCli = &r
}
