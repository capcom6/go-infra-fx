package http

import (
	"github.com/capcom6/go-infra-fx/cli"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"http",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("http")
	}),
	fx.Provide(
		New,
	),
)

func init() {
	cli.Register("http:run", Run)
}
