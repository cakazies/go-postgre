package api

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func SentryInit(err error) {
	sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry.dsn"),
	})

	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
}
