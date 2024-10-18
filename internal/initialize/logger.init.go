package initialize

import (
	"github.com/longln/reboot-simplebank/global"
	"github.com/longln/reboot-simplebank/pkg/logger"
)


func InitLogger() {
	logger := logger.NewLogger(global.Config.LogConfig)
	global.Logger = logger
	global.Logger.Info("init logger success")
}