package fxutil_test

import (
	"context"
	"testing"

	"github.com/capcom6/go-infra-fx/fxutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func TestWithNamedLogger(t *testing.T) {
	var logger *zap.Logger

	// Create fx.App with zap.NewNop for testing
	app := fx.New(
		fx.Provide(zap.NewNop),
		fxutil.WithNamedLogger("expectedName"),
		fx.Populate(&logger),
	)

	// Start the app
	require.NoError(t, app.Start(context.TODO()))

	// Ensure cleanup happens even if test fails
	t.Cleanup(func() {
		require.NoError(t, app.Stop(context.TODO()))
	})

	// Assert that the logger has the expected name
	require.Equal(t, "expectedName", logger.Name())
}
