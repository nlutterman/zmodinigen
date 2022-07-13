package steamapi

import (
	"encoding/json"
	"gitlab.com/nlutterman/zmodinigen/config"
	"gitlab.com/nlutterman/zmodinigen/errors"
	"gitlab.com/nlutterman/zmodinigen/steamworkshop"
	"io"
)

type Client struct {
	*config.Config
}

// NewSteamAPIClient
// TODO: Make sure we rate limit the requests the SteamAPIClient makes
func NewSteamAPIClient(config *config.Config) *Client {
	return &Client{
		Config: config,
	}
}

// GetCollections queries the Steam API for the provided collection IDs
func (client *Client) GetCollections(collectionIDs []string) (steamworkshop.Collections, *errors.AppError) {
	var collections = make(steamworkshop.Collections)
	var appErr *errors.AppError

	if len(collectionIDs) == 0 {
		return collections, appErr
	}

	request := NewWorkshopCollectionRequest(collectionIDs)
	response, appErr := request.Exec(client.SteamAPIEndpoints)
	if appErr != nil {
		return collections, appErr
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return collections, errors.NewError(errors.ErrorInternal, "error reading body of HTTP response: %v", err)
	}

	err = json.Unmarshal(body, &collections)
	if err != nil {
		return collections, errors.NewError(errors.ErrorInternal, "error unmarshalling response from Stream API: %v", err)
	}

	return collections, nil
}

func (client *Client) GetCollectionItemData(collection []steamworkshop.Item) map[string][]steamworkshop.Item {
	var itemsByCollectionID map[string][]steamworkshop.Item
	if len(collection) == 0 {
		return itemsByCollectionID
	}

	return itemsByCollectionID
}
