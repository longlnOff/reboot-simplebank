package logger

import (
	"os"

	"github.com/longln/reboot-simplebank/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LogConfig) *LoggerZap {
	loglevel := config.Level
	var level zapcore.Level

	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoder := setupLogFormat()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout)), 
		level)
	
		return &LoggerZap{
			Logger: zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		}
}

func setupLogFormat() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// config time format
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// config text
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// config Caller
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// config time key
	encodeConfig.TimeKey = "time"
	return zapcore.NewConsoleEncoder(encodeConfig)
}