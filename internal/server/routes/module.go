package routes

import (
	"github.com/sensfo/server/domain"
	"go.uber.org/fx"
)

func Module() fx.Option {
	createEntityRoute := fx.Annotate(
		NewCreateRoute,
		fx.As(new(domain.Route)),
		fx.ResultTags(`group:"routes"`),
	)

	return fx.Module("routes", fx.Provide(createEntityRoute))
}
