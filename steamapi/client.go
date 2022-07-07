package steamapi

import (
	"fmt"
	"github.com/nlutterman/zmodinigen/config"
	"github.com/nlutterman/zmodinigen/errors"
	"github.com/nlutterman/zmodinigen/steamworkshop"
	"io"
	"net/http"
)

type Client struct {
	*config.Config
}

func NewSteamAPIClient(config *config.Config) *Client {
	return &Client{
		Config: config,
	}
}

func (client *Client) GetCollectionInfo(collectionIDs []string) ([]steamworkshop.Item, *errors.AppError) {
	var collections []steamworkshop.Item
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
		return collections, errors.NewError(errors.ErrorInternal, "error reading body of HTTP response: %v", appErr)
	}

	response.

	return collections
}

func (client *Client) GetCollectionItemData(collection []steamworkshop.Item) map[string][]steamworkshop.Item {
	var itemsByCollectionID map[string][]steamworkshop.Item
	if len(collection) == 0 {
		return itemsByCollectionID
	}

	return itemsByCollectionID
}
