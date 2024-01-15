package controller

import (
	"AlarmPawServer/database"
	"AlarmPawServer/push"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResp{
		Code:      200,
		Message:   "pong",
		Timestamp: time.Now().Unix(),
	})
}

func BaseController(c *gin.Context) {

	params, err := push.ToParamsHandler(c)
	if err != nil {
		c.JSON(http.StatusOK, failed(400, "failed to get device token: %v", err))
		return
	}
	err = push.Push(params)

	if err != nil {
		c.JSON(http.StatusOK, failed(500, "push failed: %v", err))
		return
	}

	c.JSON(http.StatusOK, success())

}

func GetInfo(c *gin.Context) {

}

func RegisterController(c *gin.Context) {
	var deviceKey, deviceToken string

	deviceToken = c.Param("deviceToken")

	if c.Request.Method == "POST" {
		deviceKey = c.PostForm("deviceKey")
		deviceToken = c.PostForm("deviceToken")
	}

	if deviceToken == "" {
		c.JSON(http.StatusOK, failed(400, "deviceToken is empty"))
		return
	}

	newKey, err := database.DB.SaveDeviceTokenByKey(deviceKey, deviceToken)
	if err != nil {
		c.JSON(http.StatusOK, failed(500, "device registration failed: %v", err))
	}

	c.JSON(http.StatusOK, data(map[string]string{
		"key":          newKey,
		"device_key":   newKey,
		"device_token": deviceToken,
	}))
}
