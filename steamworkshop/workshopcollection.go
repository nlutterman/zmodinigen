package steamworkshop

type Collection struct {
	Publishedfileid string `json:"publishedfileid"`
	Result          int    `json:"result"`
	Children        []struct {
		Publishedfileid string `json:"publishedfileid"`
		Sortorder       int    `json:"sortorder"`
		Filetype        int    `json:"filetype"`
	} `json:"children"`
}
