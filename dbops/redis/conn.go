package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	rdbConn *redis.Client
)

//Init ...
func Init() (err error) {
	rdbConn = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
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
