package engine

import (
	"github.com/sensfo/server/domain"
	"github.com/sensfo/server/internal/server/routes"
	"go.uber.org/fx"
)

func Module() fx.Option {
	annotatedEngine := fx.Annotate(NewGinEngine, fx.As(new(domain.Engine)), fx.ParamTags(`group:"routes"`))

	return fx.Module("server-engine", routes.Module(), fx.Provide(annotatedEngine))
}
