package controller

import (
	"AlarmPawServer/config"
	"AlarmPawServer/database"
	"AlarmPawServer/push"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"strings"
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

	params, err := ToParamsHandler(c)
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
	devices, _ := database.DB.CountAll()
	c.JSON(200, map[string]interface{}{
		"version": "1.0.0",
		"build":   "",
		"arch":    runtime.GOOS + "/" + runtime.GOARCH,
		"commit":  "",
		"devices": devices,
	})

}

func RegisterController(c *gin.Context) {
	var deviceKey, deviceToken string

	for _, v := range c.Params {
		if strings.ToLower(v.Key) == config.DeviceKey {
			deviceKey = v.Value
		} else if strings.ToLower(v.Key) == config.DeviceToken {
			deviceToken = v.Value
		}
	}

	for k, v := range c.Request.URL.Query() {
		if strings.ToLower(k) == config.DeviceKey && deviceKey == "" {
			deviceKey = v[0]
		} else if strings.ToLower(k) == config.DeviceToken && deviceToken == "" {
			deviceToken = v[0]
		}
	}

	if c.Request.Method == "POST" {
		for k, v := range c.Request.PostForm {
			if strings.ToLower(k) == config.DeviceKey && deviceKey == "" {
				deviceKey = v[0]
			} else if strings.ToLower(k) == config.DeviceToken && deviceToken == "" {
				deviceToken = v[0]
			}
		}
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
