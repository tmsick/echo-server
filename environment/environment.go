package environment

import (
	"github.com/kelseyhightower/envconfig"
)

const prefix = "ECHO_SERVER"

var (
	Env       string
	JWTSecret string
)

func init() {
	e, err := Load()
	if err != nil {
		panic(err)
	}

	Env = e.Env
	JWTSecret = e.JWTSecret
}

type Environment struct {
	Env       string `envconfig:"ECHO_SERVER_ENV"`
	JWTSecret string `envconfig:"ECHO_SERVER_JWT_SECRET"`
}

func Load() (*Environment, error) {
	e := new(Environment)
	if err := envconfig.Process(prefix, e); err != nil {
		return nil, err
	}
	return e, nil
}
