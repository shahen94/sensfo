package data

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	annotated := fx.Annotate(NewEntityData, fx.As(new(Repository)))

	return fx.Module("data", fx.Provide(annotated), fx.Provide(NewDataSource))
}
