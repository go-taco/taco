package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
)

func LoadEnvs[T any]() (configs T, err error) {
	err = cleanenv.ReadConfig(".local.env", &configs)

	return
}
