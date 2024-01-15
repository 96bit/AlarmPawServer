package main

import (
	"AlarmPawServer/config"
	"AlarmPawServer/controller"
	"AlarmPawServer/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(config.LocalConfig.System.Mode)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	router.GET("/ping", controller.Ping)
	router.GET("/register/:deviceToken", controller.RegisterController)
	router.POST("/register", controller.RegisterController)
	router.GET("/info", controller.GetInfo)
	router.POST("/push", controller.BaseController)
	router.GET("/:device_key", controller.BaseController)
	router.GET("/:device_key/:params1", controller.BaseController)
	router.GET("/:device_key/:params1/:params2", controller.BaseController)
	router.GET("/:device_key/:params1/:params2/:params3", controller.BaseController)
	router.POST("/:device_key", controller.BaseController)
	router.POST("/:device_key/:params1", controller.BaseController)
	router.POST("/:device_key/:params1/:params2", controller.BaseController)
	router.POST("/:device_key/:params1/:params2/:params3", controller.BaseController)
	addr := config.LocalConfig.System.Host + ":" + config.LocalConfig.System.Post
	if err := router.Run(addr); err != nil {
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
