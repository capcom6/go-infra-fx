package db

import (
	"github.com/capcom6/go-infra-fx/cli"
	"github.com/capcom6/go-infra-fx/fxutil"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"db",
	fxutil.WithNamedLogger("db"),
	fx.Provide(
		New,
		NewSQL,
	),
)

func init() {
	cli.Register("db:auto-migrate", AutoMigrate)
	cli.Register("db:migrate", Migrate)
}
