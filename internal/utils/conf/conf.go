package conf
import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() {
	loadConfig()
	watchConfig()
	log.Println("Config Load Completed.")
}

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		log.Printf("Config changed: %s.\n", e.Name)
	})
}
