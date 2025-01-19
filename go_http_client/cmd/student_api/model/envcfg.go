package model

type EnvCfg struct {
	Endpoint   string `envconfig:"ENDPOINT" required:"true"`
	AuthHeader string `envconfig:"AUTH_HEADER" required:"true"`
}
