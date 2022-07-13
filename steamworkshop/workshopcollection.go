package steamworkshop

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

type Collections map[string]Collection

func (collection Collections) String() string {
	collectionJson, err := json.Marshal(collection)
	if err != nil {
		log.Err(err).Msg("error marshalling Collections collection")
	}
	return string(collectionJson)
}

type Collection struct {
	PublishedFileID string `json:"publishedfileid"`
	Result          int    `json:"result"`
	Children        []struct {
		PublishedFileID string `json:"publishedfileid"`
		SortOrder       int    `json:"sortorder"`
		Filetype        int    `json:"filetype"`
	} `json:"children"`
}
