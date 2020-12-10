package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger to be global
var Logger = getLogger()

// Get logger
func getLogger() *logrus.Logger {
	logger := logrus.New()

	switch os.Getenv("APP_LOG_LEVEL") {
	case "debug", "5":
		logger.SetLevel(logrus.DebugLevel)
	case "info", "4":
		logger.SetLevel(logrus.WarnLevel)
	case "warn", "3":
		logger.SetLevel(logrus.WarnLevel)
	case "error", "2":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal", "1":
		logger.SetLevel(logrus.FatalLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	return logger
}
