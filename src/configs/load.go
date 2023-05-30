package configs

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

func LoadEnvs[T any]() (configs T, err error) {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "local"
	}

	path := os.Getenv("CONFIG_FILE_PATH")

	filePath := fmt.Sprintf("%s.%s.env", path, mode)

	err = cleanenv.ReadConfig(filePath, &configs)

	return
}
