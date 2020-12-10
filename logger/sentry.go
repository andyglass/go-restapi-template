package logger

import (
	"time"

	"github.com/onrik/logrus/sentry"
	"github.com/sirupsen/logrus"
)

// InitSentry hook
func InitSentry(sentryDSN, version string) error {
	if sentryDSN == "" {
		return nil
	}
	log := Logger
	opts := sentry.Options{Dsn: sentryDSN}
	hook, err := sentry.NewHook(
		opts,
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	)
	if err != nil {
		return err
	}

	hook.AddTag("project", "go-restapi")
	hook.SetRelease("1.0.0")
	hook.SetEnvironment("dev")
	hook.SetFlushTimeout(1 * time.Second)

	log.AddHook(hook)
	log.Info("Sentry hook initialized successfully")

	return nil
}
