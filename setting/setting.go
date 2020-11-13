package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Init ...
//使用viper库管理配置
func Init() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	err = viper.ReadInConfig()
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return
}
