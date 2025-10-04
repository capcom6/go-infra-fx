package fxutil

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// WithNamedLogger returns an fx.Option that decorates the zap.Logger
// with a named logger. This is useful for distinguishing log output
// from different modules.
func WithNamedLogger(name string) fx.Option {
	return fx.Decorate(func(log *zap.Logger) *zap.Logger { return log.Named(name) })
}
