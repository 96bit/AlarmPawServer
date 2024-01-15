package push

import (
	"AlarmPawServer/config"
	"AlarmPawServer/database"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"strings"
	"time"
)

func Push(params map[string]string) error {

	pl := payload.NewPayload().
		AlertTitle(VerifyMap(params, Title)).
		AlertBody(VerifyMap(params, Body)).
		Sound(VerifyMap(params, Sound)).
		Category(VerifyMap(params, Category))

	for k, v := range params {

		if k == Group && VerifyMap(params, Group) != "" {
			pl = pl.ThreadID(params[Group])
			continue
		}

		if k == DeviceKey || k == DeviceToken || k == Title || k == Body || k == Sound || k == Category {
			continue
		}

		pl.Custom(strings.ToLower(k), v)

	}

	resp, err := CLI.Push(&apns2.Notification{
		DeviceToken: params[DeviceToken],
		Topic:       config.LocalConfig.Apple.Topic,
		Payload:     pl.MutableContent(),
		Expiration:  time.Now().Add(24 * time.Hour),
	})
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("APNS push failed: %s", resp.Reason)
	}
	return nil
}

func ToParamsHandler(c *gin.Context) (map[string]string, error) {
	var err error
	var paramsResult = make(map[string]string)
	// 获取所有路径参数
	switch len(c.Params) {

	case 1:
		paramsResult[DeviceKey] = c.Params[0].Value
	case 2:
		paramsResult[DeviceKey] = c.Params[0].Value
		paramsResult[Body] = c.Params[1].Value
	case 3:
		paramsResult[DeviceKey] = c.Params[0].Value
		paramsResult[Title] = c.Params[1].Value
		paramsResult[Body] = c.Params[2].Value
	case 4:
		paramsResult[DeviceKey] = c.Params[0].Value
		paramsResult[Category] = c.Params[1].Value
		paramsResult[Title] = c.Params[2].Value
		paramsResult[Body] = c.Params[3].Value

	}

	// 获取所有url参数
	var params = c.Request.URL.Query()

	for k, v := range params {
		if value, ok := paramsResult[k]; !ok {
			paramsResult[k] = v[0]
		} else if value == "" {
			paramsResult[k] = v[0]
		}
	}

	// 获取所有post参数
	if c.Request.Method == "POST" {
		var postParams = c.Request.PostForm

		for k, v := range postParams {
			if value, ok := paramsResult[k]; !ok {
				paramsResult[k] = v[0]
			} else if value == "" {
				paramsResult[k] = v[0]
			}

		}
	}

	// 处理默认值
	if VerifyMap(paramsResult, IsArchive) == "" {
		paramsResult[IsArchive] = config.IsArchive
	}
	if VerifyMap(paramsResult, AutoCopy) == "" {
		paramsResult[AutoCopy] = config.AutoCopy
	}
	if VerifyMap(paramsResult, Level) == "" {
		paramsResult[Level] = config.LevelA
	}
	if VerifyMap(paramsResult, Category) == "" {
		paramsResult[Category] = config.CategoryDefault
	}

	if VerifyMap(paramsResult, Sound) != "" && !strings.HasSuffix(paramsResult[Sound], ".caf") {
		paramsResult[Sound] = paramsResult[Sound] + ".caf"
	}

	if VerifyMap(paramsResult, DeviceToken) == "" {
		if VerifyMap(paramsResult, DeviceKey) == "" {
			err = errors.New("deviceKey or deviceToken is required")
			return nil, err
		}
		paramsResult[DeviceToken], err = database.DB.DeviceTokenByKey(paramsResult[DeviceKey])
		if err != nil {
			err = errors.New("failed to get device token: " + err.Error())
			return nil, err
		}
	}

	return paramsResult, nil
}

func VerifyMap(data map[string]string, key string) string {
	if value, ok := data[key]; ok {
		return value
	}
	return ""
}
