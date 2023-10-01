package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenAiApiKey string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading env: %s", err)
	}

	conf := Config{}
	conf.OpenAiApiKey = os.Getenv("OPENAI_API_KEY")
	return conf
}
