package controller

import (
	"AlarmPawServer/config"
	"AlarmPawServer/database"
	"AlarmPawServer/modal"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ParamsHandler(c *gin.Context) (result modal.Params, err error) {

	// 获取所有url参数
	switch len(c.Params) {

	case 1:
		result.DeviceKey = c.Params[0].Value
	case 2:
		result.DeviceKey = c.Params[0].Value
		result.Body = c.Params[1].Value
	case 3:
		result.DeviceKey = c.Params[0].Value
		result.Title = c.Params[1].Value
		result.Body = c.Params[2].Value
	case 4:
		result.DeviceKey = c.Params[0].Value
		result.Category = c.Params[1].Value
		result.Title = c.Params[2].Value
		result.Body = c.Params[3].Value

	}
	var params = c.Request.URL.Query()

	if result.DeviceKey == "" && len(params["deviceKey"]) > 0 {
		result.DeviceKey = params["deviceKey"][0]
	}

	if result.DeviceToken == "" && len(params["deviceToken"]) > 0 {
		result.DeviceToken = params["deviceToken"][0]
	}

	if result.Title == "" && len(params["title"]) > 0 {
		result.Title = params["title"][0]
	}

	if result.Body == "" && len(params["body"]) > 0 {
		result.Body = params["body"][0]
	}

	if result.Category == "" && len(params["category"]) > 0 {
		result.Category = params["category"][0]
	}

	if len(params["icon"]) > 0 {
		result.Icon = params["icon"][0]
	}
	if len(params["image"]) > 0 {
		result.Image = params["image"][0]
	}
	if len(params["url"]) > 0 {
		result.Url = params["url"][0]
	}
	if len(params["isArchive"]) > 0 {
		result.IsArchive = params["isArchive"][0]
	}
	if len(params["group"]) > 0 {
		result.Group = params["group"][0]
	}
	if len(params["sound"]) > 0 {
		result.Sound = params["sound"][0]
	}
	if len(params["autoCopy"]) > 0 {
		result.AutoCopy = params["autoCopy"][0]
	}
	if len(params["copy"]) > 0 {
		result.Copy = params["copy"][0]
	}
	if len(params["badge"]) > 0 {
		result.Badge = params["badge"][0]
	}

	if len(params["level"]) > 0 {
		result.Level = params["level"][0]
	}

	if len(params["cipherText"]) > 0 {
		result.CipherText = params["cipherText"][0]
	}

	if c.Request.Method == "POST" {
		var postParams = c.Request.PostForm

		if result.DeviceKey == "" && len(postParams["deviceKey"]) != 0 {
			result.DeviceKey = postParams["deviceKey"][0]
		}

		if result.DeviceToken == "" && len(postParams["deviceToken"]) != 0 {
			result.DeviceToken = postParams["deviceToken"][0]
		}

		if result.Title == "" && len(postParams["title"]) != 0 {
			result.Title = postParams["title"][0]
		}
		if result.Body == "" && len(postParams["body"]) != 0 {
			result.Body = postParams["body"][0]
		}

		if result.Category == "" && len(postParams["category"]) != 0 {
			result.Category = postParams["category"][0]
		}
		if result.Icon == "" && len(postParams["icon"]) > 0 {
			result.Icon = postParams["icon"][0]
		}
		if result.Image == "" && len(postParams["image"]) > 0 {
			result.Image = postParams["image"][0]
		}
		if result.Url == "" && len(postParams["url"]) > 0 {
			result.Url = postParams["url"][0]
		}
		if result.IsArchive == "" && len(postParams["isArchive"]) > 0 {
			result.IsArchive = postParams["isArchive"][0]
		}
		if result.Group == "" && len(postParams["group"]) > 0 {
			result.Group = postParams["group"][0]
		}
		if result.Sound == "" && len(postParams["sound"]) > 0 {
			result.Sound = postParams["sound"][0]
		}
		if result.AutoCopy == "" && len(postParams["autoCopy"]) > 0 {
			result.AutoCopy = "1"
		}
		if result.Copy == "" && len(postParams["copy"]) > 0 {
			result.Copy = postParams["copy"][0]
		}
		if result.Badge == "" && len(postParams["badge"]) > 0 {
			result.Badge = postParams["badge"][0]
		}
		if result.Level == "" && len(postParams["level"]) > 0 {
			result.Level = postParams["level"][0]
		}
		if result.CipherText == "" && len(postParams["cipherText"]) > 0 {
			result.CipherText = postParams["cipherText"][0]
		}
	}

	// 处理默认值
	if result.IsArchive == "" {
		result.IsArchive = config.IsArchive
	}
	if result.AutoCopy == "" {
		result.AutoCopy = config.AutoCopy
	}
	if result.Level == "" {
		result.Level = config.LevelA
	}
	if result.Category == "" {
		result.Category = config.CategoryDefault
	}

	if result.Sound != "" {
		if !strings.HasSuffix(result.Sound, ".caf") {
			result.Sound = result.Sound + ".caf"
		}
	}

	if result.DeviceToken == "" {
		if result.DeviceKey == "" {
			err = errors.New("deviceKey or deviceToken is required")
			c.JSON(200, failed(400, "deviceKey or deviceToken is required"))
			return
		}
		result.DeviceToken, err = database.DB.DeviceTokenByKey(result.DeviceKey)
		if err != nil {
			c.JSON(200, failed(400, "failed to get device token: %v", err))
			return
		}
	}

	return
}
