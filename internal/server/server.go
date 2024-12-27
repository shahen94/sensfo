package server

import (
	"context"

	"github.com/sensfo/server/domain"
	"github.com/sensfo/server/internal/configuration"
	"go.uber.org/fx"
)

type Server struct {
	engine        domain.Engine
	configuration *configuration.Configuration
}

func (e *Server) Start(context.Context) error {
	err := e.engine.ListenAndServe(e.configuration.Port)

	return err
}

func (e *Server) Stop(context.Context) error {
	return e.engine.Shutdown()
}

// New creates a new Engine instance.
func New(engine domain.Engine, configuration *configuration.Configuration, lc fx.Lifecycle) *Server {
	server := &Server{
		engine:        engine,
		configuration: configuration,
	}

	lc.Append(fx.Hook{
		OnStart: server.Start,
		OnStop:  server.Stop,
	})

	return server
}
