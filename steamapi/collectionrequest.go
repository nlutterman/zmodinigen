package steamapi

import (
	"github.com/nlutterman/zmodinigen/errors"
	"github.com/nlutterman/zmodinigen/steamworkshop"
	"net/http"
	"net/url"
	"strconv"
)

const collectionCountKey = "collectioncount"
const collectionIDKey = "publishedfileids"

type WorkshopCollectionRequest struct {
	*Request
	CollectionIDs []string
}

type WorkshopCollectionResponse struct {
	// TODO: Check on pagination
	Response struct {
		Result      int                        `json:"result"`
		ResultCount int                        `json:"resultcount"`
		Collections []steamworkshop.Collection `json:"collectiondetails"`
	} `json:"response"`
}

func (cr *WorkshopCollectionRequest) GetRequestData() url.Values {
	if cr.data == nil {
		cr.data = url.Values{
			collectionCountKey: []string{strconv.Itoa(len(cr.CollectionIDs))},
			collectionIDKey:    cr.CollectionIDs,
		}
	}
	return cr.data
}

func (cr *WorkshopCollectionRequest) Exec(endpoints EndpointMap) (*http.Response, *errors.AppError) {
	return cr.Request.exec(http.MethodPost, endpoints.GetURL(cr.endpoint), cr.GetRequestData())
}

func NewWorkshopCollectionRequest(ids []string) *WorkshopCollectionRequest {
	return &WorkshopCollectionRequest{
		&Request{
			endpoint: WorkshopCollectionEndpointID,
		}, ids,
	}
}
