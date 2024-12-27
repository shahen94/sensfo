package configuration

import "github.com/caarlos0/env"

const (
	envDevelopment = "development"
	envProduction  = "production"
)

type Configuration struct {
	Port int    `env:"PORT" envDefault:"8080"`
	Env  string `env:"ENV" envDefault:"development"`
}

func (c *Configuration) IsDevelopment() bool {
	return c.Env == envDevelopment
}

func (c *Configuration) IsProduction() bool {
	return c.Env == envProduction
}

func New() *Configuration {
	cfg := &Configuration{}

	err := env.Parse(cfg)

	if err != nil {
		panic(err)
	}

	return cfg
}
