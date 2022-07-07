package steamapi

type EndpointID int

type EndpointMap map[EndpointID]string

func (endpoints EndpointMap) GetURL(id EndpointID) string { return endpoints[id] }

const (
	WorkshopCollectionEndpointID EndpointID = iota
	WorkshopItemEndpointID
)

const (
	DefaultHost                  = "api.steampowered.com"
	DefaultWorkshopCollectionURI = "/ISteamRemoteStorage/GetCollectionDetails/v1/"
	DefaultWorkshopItemURI       = "/ISteamRemoteStorage/GetPublishedFileDetails/v1/"
)

var DefaultURLs = EndpointMap{
	WorkshopCollectionEndpointID: "https://" + DefaultHost + DefaultWorkshopCollectionURI,
	WorkshopItemEndpointID:       "https://" + DefaultHost + DefaultWorkshopItemURI,
}
