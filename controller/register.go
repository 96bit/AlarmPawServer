package controller

import (
	"AlarmPawServer/database"
	"github.com/gin-gonic/gin"
)

func RegisterController(c *gin.Context) {
	var deviceKey, deviceToken string

	deviceToken = c.Param("deviceToken")

	if c.Request.Method == "POST" {
		deviceKey = c.PostForm("deviceKey")
		deviceToken = c.PostForm("deviceToken")
	}

	if deviceToken == "" {
		c.JSON(200, failed(400, "deviceToken is empty"))
		return
	}

	newKey, err := database.DB.SaveDeviceTokenByKey(deviceKey, deviceToken)
	if err != nil {
		c.JSON(200, failed(500, "device registration failed: %v", err))
	}

	c.JSON(200, data(map[string]string{
		"key":          newKey,
		"device_key":   newKey,
		"device_token": deviceToken,
	}))
}
