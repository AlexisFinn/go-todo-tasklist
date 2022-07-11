package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoUri() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	return os.Getenv("MONGOURI")
}
