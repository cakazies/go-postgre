package api

import (
	"time"

	cf "github.com/cakazies/go-postgre/application/models"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func init() {
	err := cf.Connect()
	if err != nil {
		SentryInit(err)
	}
}

func SentryInit(err error) {
	sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry.dsn"),
	})

	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
}
