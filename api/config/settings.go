package config

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"

	"os"
)

type CTXUserID string

const (
	AuthKey              = "KEY"
	Connection           = "CONN"
	ClientURI            = "CLIENT"
	UserIDKey  CTXUserID = "userId"
)

type Settings struct {
	AuthKey    string
	ConnString string
	ClientURI  string
}

func GetSettings() Settings {
	mode := Mode()
	if mode == "prod" {
		return settingsFromVault()
	}

	return settingsFromEnv()
}

func settingsFromEnv() Settings {
	return Settings{
		AuthKey:    getEnvVar(AuthKey),
		ConnString: getEnvVar(Connection),
		ClientURI:  getEnvVar(ClientURI),
	}
}

func settingsFromVault() Settings {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	uri := getEnvVar("VAULT_URI")
	client, err := azsecrets.NewClient(uri, cred, nil)
	if err != nil {
		log.Fatalf("failed to create vault client: %v", err)
	}

	settings := Settings{}
	version := ""
	authKeyResp, err := client.GetSecret(context.TODO(), AuthKey, version, nil)
	if err != nil {
		log.Fatalf("failed to get auth key: %v", err)
	}
	settings.AuthKey = *authKeyResp.Value

	connResp, err := client.GetSecret(context.TODO(), Connection, version, nil)
	if err != nil {
		log.Fatalf("failed to get connection: %v", err)
	}
	settings.ConnString = *connResp.Value
	settings.ClientURI = "dm-api"
	return settings
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
