package push

import (
	"AlarmPawServer/config"
	"AlarmPawServer/modal"
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"time"
)

var (
	CLI *apns2.Client
)

func Push(params modal.Params) error {
	pl := payload.NewPayload().
		AlertTitle(params.Title).
		AlertBody(params.Body).
		Sound(params.Sound).
		Category(params.Category)

	if params.Group != "" {
		pl = pl.ThreadID(params.Group)
	}

	if params.Icon != "" {
		pl = pl.Custom("icon", params.Icon)
	}
	if params.Image != "" {
		pl = pl.Custom("image", params.Image)
	}
	if params.Url != "" {
		pl = pl.Custom("url", params.Url)
	}
	if params.IsArchive != "" {
		pl = pl.Custom("isArchive", params.IsArchive)
	}
	if params.AutoCopy != "" {
		pl = pl.Custom("autoCopy", params.AutoCopy)
	}
	if params.Copy != "" {
		pl = pl.Custom("copy", params.Copy)
	}
	if params.Badge != "" {
		pl = pl.Custom("badge", params.Badge)
	}
	if params.Level != "" {
		pl = pl.Custom("level", params.Level)
	}
	if params.CipherText != "" {
		pl = pl.Custom("cipherText", params.CipherText)
	}

	resp, err := CLI.Push(&apns2.Notification{
		DeviceToken: params.DeviceToken,
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
