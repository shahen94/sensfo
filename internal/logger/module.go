package logger

import "go.uber.org/fx"

func Module() fx.Option {
	annotated := fx.Annotate(NewSlogLogger, fx.As(new(Logger)))

	return fx.Module("logger", fx.Provide(annotated))
}
