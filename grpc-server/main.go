package main

import (
	"github.com/eminoz/todo/pkg/config"
	"github.com/eminoz/todo/pkg/database"
	"github.com/eminoz/todo/pkg/rounter"
)

func main() {
	config.SetupConfig()
	database.SetDatabase()
	rounter.BaseRpc()
}
