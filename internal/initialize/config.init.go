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
	os.Setenv("CONFIG_PATH", "/home/longln/SourceCode/github.com/longln/reboot-simplebank/local")
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