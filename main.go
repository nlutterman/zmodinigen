package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/nlutterman/zmodinigen/config"
	"gitlab.com/nlutterman/zmodinigen/steamapi"
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
		Msg("querying for steam workshop collection data")
	collections, err := api.GetCollections(collectionIDs)
	if err != nil {
		log.Err(err).Msg("error getting collection info")
	}

	log.Info().
		Stringer("collections", collections).
		Int("CollectionCount", len(collectionIDs)).
		Msg("querying for steam workshop item data")
	itemsByCollectionID := api.GetCollectionItemData(collections)

	return fmt.Sprintf("%v\n", itemsByCollectionID)
}
