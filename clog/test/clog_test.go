package test

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/wano/contextlog/clog"
	"testing"
)

func TestGlobal(t *testing.T) {

	clog.Debug(`@@@@`)
	clog.SetGlobalLevel(zerolog.ErrorLevel)
	clog.Info(`@@@@F`)
}

func TestContext(t *testing.T) {

	clog.SetGlobalLevel(zerolog.InfoLevel)
	ctx := context.Background()
	logger := clog.NewContextLogger()
	logger.SetPrefix(`app_name`, `app_server`)
	logger.SetPrefix(`context_id`, `random_value`)
	ctx = logger.WithContext(ctx)

	clog.Ctx(ctx).Infof(`message %s`, `!!`)

}
