package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Conf ...
//全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

//AppConfig ...
type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

//LogConfig ...
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

//MySQLConfig ...
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"use"`
	Password     string `mapstructure:"password"`
	DBname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

//RedisConfig ...
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

//Init ...
//使用viper库管理配置
func Init() (err error) {
	viper.SetConfigFile("config.json")
	// viper.SetConfigName("config")
	// viper.SetConfigType("json")
	viper.AddConfigPath("../")

	err = viper.ReadInConfig()
	if err != nil { //读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return err
	}

	//把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	//监控配置文件的修改
	viper.WatchConfig()
	viper.OnConfigChange(
		func(in fsnotify.Event) {
			fmt.Println("配置文件修改了...")
			if err := viper.Unmarshal(Conf); err != nil {
				fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
			}
		})

	return nil
}
