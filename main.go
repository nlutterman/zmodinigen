package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/nlutterman/zmodinigen/config"
	"gitlab.com/nlutterman/zmodinigen/steamapi"
	"gitlab.com/nlutterman/zmodinigen/utils"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	c := config.NewConfig()
	api := steamapi.NewSteamAPIClient(c.ClientConfig)

	collectionIDs := []string{
		"2757817552", // Squirtle Squad
	}

	log.Print(GenerateINI(api, collectionIDs))
}

func GenerateINI(api *steamapi.Client, collectionIDs []string) string {
	log.Info().Msg("starting to generate INI file(s)")

	if api == nil {
		// TODO: set up proper errors and error handling
		panic("nil api provided, provide an instantiated SteamAPIClient struct")
	}

	log.Info().
		Strs("CollectionIDs", collectionIDs).
		Int("CollectionCount", len(collectionIDs)).
		Msg("querying for steam workshop collection and item data")

	idSet := utils.NewSetFromSlice[string](collectionIDs)
	collections, err := api.GetCollections(idSet)
	if err != nil {
		log.Err(err).Msg("error getting collection info")
	}

	return fmt.Sprintf("%v\n", collections)
}
