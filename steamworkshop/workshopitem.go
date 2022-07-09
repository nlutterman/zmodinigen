package steamworkshop

import (
	"encoding/json"
)

type Item struct {
	ID                    string `json:"publishedfileid"`
	Result                int    `json:"result"`
	Creator               string `json:"creator"`
	CreatorAppID          int    `json:"creator_app_id"`
	ConsumerAppID         int    `json:"consumer_app_id"`
	Filename              string `json:"filename"`
	FileSize              int    `json:"file_size"`
	FileURL               string `json:"file_url"`
	PreviewURL            string `json:"preview_url"`
	HcontentFile          string `json:"hcontent_file"` // These are UGC files, they contain the path to the actual mod content, I think
	HcontentPreview       string `json:"hcontent_preview"`
	Title                 string `json:"title"`
	Description           string `json:"description"`
	TimeCreated           int    `json:"time_created"`
	TimeUpdated           int    `json:"time_updated"`
	Visibility            int    `json:"visibility"`
	Banned                int    `json:"banned"`
	BanReason             string `json:"ban_reason"`
	Subscriptions         int    `json:"subscriptions"`
	Favorited             int    `json:"favorited"`
	LifetimeSubscriptions int    `json:"lifetime_subscriptions"`
	LifetimeFavorited     int    `json:"lifetime_favorited"`
	Views                 int    `json:"views"`
	Tags                  []struct {
		Tag string `json:"tag"`
	} `json:"tags"`
}

func (i []Item) String() string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}
