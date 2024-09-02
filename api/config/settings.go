package config

import (
	"log"

	"os"
)

type CTXUserID string

const (
	UserIDKey CTXUserID = "userId"
)

func Key() string {
	return getEnvVar("KEY")
}

func PGConn() string {
	return getEnvVar("CONN")
}

func Mode() string {
	return getEnvVar("MODE")
}

func getEnvVar(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s doesn't exist or is not set", key)
	}
	return os.Getenv(key)
}
