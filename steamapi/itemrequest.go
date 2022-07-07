package steamapi

import (
	"github.com/nlutterman/zmodinigen/errors"
	"github.com/nlutterman/zmodinigen/steamworkshop"
	"net/http"
	"net/url"
	"strconv"
)

const itemCountKey = "itemcount"
const itemIDKey = "publishedfileids"

type WorkshopItemRequest struct {
	*Request
	ItemIDs []string
}

type WorkshopItemResponse struct {
	// TODO: Check on pagination
	Response struct {
		Result               int                  `json:"result"`
		ResultCount          int                  `json:"resultcount"`
		PublishedFileDetails []steamworkshop.Item `json:"publishedfiledetails"`
	} `json:"response"`
}

func (ir *WorkshopItemRequest) GetRequestData() url.Values {
	if ir.data == nil {
		ir.data = url.Values{
			itemCountKey: []string{strconv.Itoa(len(ir.ItemIDs))},
			itemIDKey:    ir.ItemIDs,
		}
	}
	return ir.data
}

func (ir *WorkshopItemRequest) Exec(endpoints EndpointMap) (*http.Response, *errors.AppError) {
	return ir.Request.exec(http.MethodPost, endpoints.GetURL(ir.endpoint), ir.GetRequestData())
}

func NewWorkshopItemRequest(ids []string) *WorkshopItemRequest {
	return &WorkshopItemRequest{
		&Request{
			endpoint: WorkshopItemEndpointID,
		}, ids,
	}
}
