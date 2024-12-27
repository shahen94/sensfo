package configuration

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("configuration", fx.Provide(New))
}
