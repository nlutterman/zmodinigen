package config

import (
	"gitlab.com/nlutterman/zmodinigen/steamapi"
	"os"
)

const defaultAppID = "108600"
const defaultAppName = "Project Zomboid"

const (
	EnvSteamAPIKey  = "STEAM_API_KEY"
	EnvSteamAppID   = "STEAM_APP_ID"
	EnvSteamAppName = "STEAM_APP_NAME"
)

type Config struct {
	*steamapi.ClientConfig
}

func (c *Config) loadEnvironment() {
	appID, isSet := os.LookupEnv(EnvSteamAppID)
	if !isSet {
		appID = defaultAppID
	}
	appName, isSet := os.LookupEnv(EnvSteamAppName)
	if !isSet {
		appName = defaultAppName
	}

	apiKey, apiKeyFound := os.LookupEnv(EnvSteamAPIKey)
	if !apiKeyFound {
		panic("need to set a Steam API key")
	}

	c.SteamAppID = appID
	c.SteamAppName = appName
	c.SteamAPIKey = apiKey
	c.SteamAPIHost = steamapi.DefaultHost
	c.SteamAPIEndpoints = steamapi.DefaultURLs // TODO: Allow changing endpoints via env var if need be
}

func NewConfig() *Config {
	config := &Config{}
	config.loadEnvironment()

	return config
}
