package server

import (
	"github.com/sensfo/server/internal/server/engine"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("server", engine.Module(), fx.Provide(New))
}
