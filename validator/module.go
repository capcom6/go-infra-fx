package validator

import (
	"github.com/capcom6/go-infra-fx/fxutil"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"validator",
	fxutil.WithNamedLogger("validator"),
	fx.Provide(New),
)
