package controller

import (
	"AlarmPawServer/config"
	"AlarmPawServer/database"
	"AlarmPawServer/push"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
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
		paramsKey := config.UnifiedParameter(v.Key)
		if paramsKey == config.DeviceKey {
			deviceKey = v.Value
		} else if paramsKey == config.DeviceToken {
			deviceToken = v.Value
		}
	}

	for k, v := range c.Request.URL.Query() {
		paramsKey := config.UnifiedParameter(k)
		if paramsKey == config.DeviceKey && deviceKey == "" {
			deviceKey = v[0]
		} else if paramsKey == config.DeviceToken && deviceToken == "" {
			deviceToken = v[0]
		}
	}

	if c.Request.Method == "POST" {
		for k, v := range c.Request.PostForm {
			paramsKey := config.UnifiedParameter(k)
			if paramsKey == config.DeviceKey && deviceKey == "" {
				deviceKey = v[0]
			} else if paramsKey == config.DeviceToken && deviceToken == "" {
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
