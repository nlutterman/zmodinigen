package steamapi

import (
	"fmt"
	config2 "github.com/nlutterman/zmodinigen/config"
	"github.com/nlutterman/zmodinigen/steamworkshop"
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
