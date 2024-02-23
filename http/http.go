package http

import (
	"time"

	"github.com/capcom6/go-infra-fx/http/jsonify"
	"github.com/capcom6/go-infra-fx/http/statuscode"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	ReadTimeout  = 5 * time.Second
	WriteTimeout = 5 * time.Second
	IdleTimeout  = 60 * time.Second
)

func New(params Params) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		ReadTimeout:           ReadTimeout,
		WriteTimeout:          WriteTimeout,
		IdleTimeout:           IdleTimeout,
		DisableStartupMessage: true,
		ErrorHandler:          errorHandler,
	})

	app.Use(recover.New())
	app.Use(fiberzap.New(fiberzap.Config{
		SkipBody: func(c *fiber.Ctx) bool {
			return c.Response().StatusCode() < 400
		},
		Logger: params.Logger,
		Fields: []string{"latency", "status", "method", "url", "ip", "ua", "body", "error"},
	}))

	for _, handler := range params.RootHandlers {
		handler.Register(app)
	}

	api := app.Group("/api")
	api.Use(jsonify.New())
	for _, handler := range params.ApiHandlers {
		handler.Register(api)
	}

	app.Use(statuscode.New())

	return app, nil
}
