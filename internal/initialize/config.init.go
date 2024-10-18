package initialize

import (
	"os"

	"github.com/longln/reboot-simplebank/global"
	"github.com/spf13/viper"
)


func LoadConfig() {
	// 1. Load config path from .env
	viper := viper.New()
	// set config folder
	viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}