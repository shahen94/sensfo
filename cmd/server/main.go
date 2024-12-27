package main

import (
	"github.com/sensfo/server/internal/configuration"
	"github.com/sensfo/server/internal/data"
	"github.com/sensfo/server/internal/encryption"
	"github.com/sensfo/server/internal/logger"
	"github.com/sensfo/server/internal/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.Module(),
		data.Module(),
		configuration.Module(),
		encryption.Module(),
		server.Module(),
		fx.Invoke(func(e *server.Server) {}),
	)

	app.Run()
}
