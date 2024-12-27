package encryption

import "go.uber.org/fx"

func Module() fx.Option {
	annotated := fx.Annotate(NewInfernoEncryption, fx.As(new(Encryption)))

	return fx.Module("encryption", fx.Provide(annotated))
}
