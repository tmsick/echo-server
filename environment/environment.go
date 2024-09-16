package environment

import (
	"github.com/kelseyhightower/envconfig"
)

const prefix = "ECHO_SERVER"

type Env struct {
	Env string
}

func Load() (*Env, error) {
	e := new(Env)
	if err := envconfig.Process(prefix, e); err != nil {
		return nil, err
	}
	return e, nil
}
