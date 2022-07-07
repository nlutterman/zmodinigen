package main

import "github.com/nlutterman/zmodinigen/steamworkshop"

type ZomboidMap struct {
	*steamworkshop.Item
	MapFolders []string
}
