package redis

import (
	"fmt"
	"github.com/Moqqll/go_webserver/setting"

	"github.com/go-redis/redis"
)

var (
	rdbConn *redis.Client
)

//Init ...
func Init(cfg *setting.RedisConfig) (err error) {
	rdbConn = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdbConn.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//Close ...
func Close() {
	_ = rdbConn.Close()
}
