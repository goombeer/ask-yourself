package di

import (
	"os"

	"github.com/goombeer/ask-yourself/backend/src/config"
)

func loadConfig() config.Config {
	switch os.Getenv("ENVIRONMENT") {
	case "dev", "staging", "production":
		return config.LoadConfig()
	default:
		return config.LoadConfigWithDotenv()
	}
}