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
	Sound            = "sound"
	AutoCopy         = "autocopy"
	Level            = "level"
)

const (
	FilePathEnv     = "ALARM_PAW_CONFIG"
	DefaultFilePath = "./config.yaml"
	TestFilePath    = "./config.test.yaml"
	ReleaseFilePath = "./config.release.yaml"
)

var (
	LocalConfig *Config
	LocalVP     *viper.Viper
)
