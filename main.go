package main

import (
	"be/configs"
	"be/utils"

	"github.com/labstack/gommon/log"
)

func main() {
	var config = configs.GetConfig()
	var db = utils.InitDB(config)

	log.Info(db)
}
