package main

import (
	"AlarmPawServer/controller"
	"AlarmPawServer/modal"
	"github.com/gin-gonic/gin"
	"time"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) { c.JSON(200, "ok") })
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, modal.CommonResp{
			Code:      200,
			Message:   "pong",
			Timestamp: time.Now().Unix(),
		})
	})
	// 设备注册
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
	return router
}
