package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

// InitLogger 初始化日志系统
func InitLogger(level string) {
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}
}
