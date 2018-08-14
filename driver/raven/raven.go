package tracing

import (
	"github.com/getsentry/raven-go"
)

func New(sentryDsn string) *raven.Client {
	if err := raven.SetDSN(sentryDsn); err != nil {
		panic(err)
	}
	return raven.DefaultClient
}
