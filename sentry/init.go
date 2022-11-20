package sentry

import (
	"fmt"

	"github.com/ario-team/glassnode-api/config"
	"github.com/getsentry/sentry-go"
)

var Sentry *sentry.Hub

func Init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Viper.GetString("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sentry initialized")
	}
	Sentry = sentry.CurrentHub().Clone()
}
