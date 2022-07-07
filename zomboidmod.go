package main

import "gitlab.com/nlutterman/zmodinigen/steamworkshop"

type ZomboidMod struct {
	*steamworkshop.Item
	ModIDs []string
}
