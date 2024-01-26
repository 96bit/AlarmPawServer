package push

import (
	"AlarmPawServer/config"
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"strings"
	"time"
)

func Push(params map[string]string) error {

	pl := payload.NewPayload().
		AlertTitle(config.VerifyMap(params, config.Title)).
		AlertBody(config.VerifyMap(params, config.Body)).
		Sound(config.VerifyMap(params, config.Sound)).
		Category(config.VerifyMap(params, config.Category))

	for k, v := range params {

		if k == config.DeviceKey ||
			k == config.DeviceToken ||
			k == config.Title ||
			k == config.Body ||
			k == config.Sound ||
			k == config.Category {
			continue
		}

		pl.Custom(strings.ToLower(k), v)

	}

	if group := config.VerifyMap(params, config.Group); group != "" {
		pl = pl.ThreadID(group)
	} else {
		pl = pl.ThreadID("default")
		params[config.Group] = "default"
	}

	resp, err := CLI.Push(&apns2.Notification{
		DeviceToken: params[config.DeviceToken],
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
