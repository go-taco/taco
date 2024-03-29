package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

func LoadEnvs[T any]() (configs T, err error) {
	env := strings.ToLower(os.Getenv("ENV"))
	if env == "" {
		env = "local"
	}

	path := os.Getenv("CONFIG_FILE_PATH")

	filePath := fmt.Sprintf("%s.%s.env", path, env)

	err = cleanenv.ReadConfig(filePath, &configs)

	return
}
