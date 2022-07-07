package main

import (
	"fmt"
	"gitlab.com/nlutterman/zmodinigen/config"
	"gitlab.com/nlutterman/zmodinigen/steamapi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	c := config.NewConfig()
	api := steamapi.NewSteamAPIClient(c)

	collectionIDs := []string{
		"2757817552", // Squirtle Squad
	}

	log.Print(GenerateINI(api, collectionIDs))
}

func GenerateINI(api *steamapi.Client, collectionIDs []string) string {
	log.Info().
		Strs("collection_ids", collectionIDs).
		Int("collection_count", len(collectionIDs)).
		Msg("creating INI file(s)")

	if api == nil {
		// TODO: set up proper errors and error handling
		panic("nil api provided, provide an instantiated SteamAPIClient struct")
	}

	log.Info().
		Strs("CollectionIDs", collectionIDs).
		Int("CollectionCount", len(collectionIDs)).
		Msg("querying for steamworkshop collection data")
	collections := api.GetCollectionInfo(collectionIDs)

	log.Info().
		Stringers("collections", collections).
		Int("CollectionCount", len(collectionIDs)).
		Msg("querying for steamworkshop item data")
	itemsByCollectionID := api.GetCollectionItemData(collections)

	return fmt.Sprintf("%v\n", itemsByCollectionID)
}
