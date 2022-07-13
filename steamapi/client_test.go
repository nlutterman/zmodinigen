package steamapi

import (
	"fmt"
	config2 "gitlab.com/nlutterman/zmodinigen/config"
	"gitlab.com/nlutterman/zmodinigen/steamworkshop"
	"testing"
)

func TestGetCollectionInfo(t *testing.T) {
	config := config2.NewConfig()
	client := NewSteamAPIClient(config)

	collection := steamworkshop.Item{
		ID: "2757817552",
	}

	fmt.Println(config, client, collection) // just to get go to shut up
}

func TestGetCollectionItemData(t *testing.T) {

}

func TestGenerateINI(t *testing.T) {

}
