package main

import "gitlab.com/nlutterman/zmodinigen/steamworkshop"

type ZomboidMap struct {
	*steamworkshop.Item
	MapFolders []string
}
