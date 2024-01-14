package main

import (
	"AlarmPawServer/config"
	"AlarmPawServer/database"
)

func main() {
	addr := config.LocalConfig.System.Host + ":" + config.LocalConfig.System.Post
	engin := Router()
	if err := engin.Run(addr); err != nil {
		panic(err)
	}

}

func init() {
	switch config.LocalConfig.System.DBType {
	case "mysql":
		database.DB = database.NewMySQL(config.GetDsn())
	default:
		database.DB = database.NewBboltdb(config.LocalConfig.System.DBPath)

	}

}
