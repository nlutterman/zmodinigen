package main

import "github.com/nlutterman/zmodinigen/steamworkshop"

type ZomboidMod struct {
	*steamworkshop.Item
	ModIDs []string
}
