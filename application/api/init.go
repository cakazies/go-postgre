package api

import (
	"errors"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func SentryInit() {
	sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry.dsn"),
	})

	sentry.CaptureException(errors.New("my error"))
	sentry.Flush(time.Second * 5)
}
