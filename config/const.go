package config

import "github.com/spf13/viper"

// URL > QUERY > POST

const (
	LevelA = "active"
	LevelT = "timeSensitive"
	LevelP = "passive"
)

const (
	CategoryDefault  = "myNotificationCategory"
	AutoCopyDefault  = "0"
	IsArchiveDefault = "1"
	DeviceKey        = "devicekey"
	DeviceToken      = "devicetoken"
	Category         = "category"
	Title            = "title"
	Body             = "body"
	IsArchive        = "isarchive"
	Group            = "group"
	DefaultGroup     = "Default"
	Sound            = "sound"
	AutoCopy         = "autocopy"
	Level            = "level"
)

const (
	FilePathEnv     = "ALARM_PAW_CONFIG"
	DefaultFilePath = "/data/config.yaml"
	TestFilePath    = "/data/config.test.yaml"
	ReleaseFilePath = "/data/config.release.yaml"
)

var (
	LocalConfig *Config
	LocalVP     *viper.Viper
)
