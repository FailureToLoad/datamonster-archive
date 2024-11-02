package config

import (
	"log"

	"os"
)

func Key() string {
	return getEnvVar("KEY")
}

func getEnvVar(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s doesn't exist or is not set", key)
	}
	return os.Getenv(key)
}
