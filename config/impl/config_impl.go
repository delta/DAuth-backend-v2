package impl

import (
	"os"

	"github.com/delta/DAuth-backend-v2/config"
	"github.com/joho/godotenv"
)

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) config.Config {
	var _ = godotenv.Load(filenames...)
	return &configImpl{}
}
