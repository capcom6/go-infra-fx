package http

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type RunServerParams struct {
	fx.In

	Server *Server
	Logger *zap.Logger
	LC     fx.Lifecycle
}

func Run(params RunServerParams) error {
	go func() {
		params.Logger.Info("Starting server...")

		if err := params.Server.Start(); err != nil {
			params.Logger.Error("Error starting server", zap.Error(err))
		}
	}()

	params.LC.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return params.Server.Stop(ctx)
		},
	})

	return nil
}
