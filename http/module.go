package http

import (
	"github.com/capcom6/go-infra-fx/cli"
	"github.com/capcom6/go-infra-fx/fxutil"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fxutil.WithNamedLogger("http"),
	fx.Provide(
		New,
		NewServer,
	),
)

func init() {
	cli.Register("http:run", Run)
}
