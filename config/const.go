package config

import "github.com/spf13/viper"

// URL > QUERY > POST

const (
	LevelA = "active"
	LevelT = "timeSensitive"
	LevelP = "passive"
)

const (
	CategoryDefault = "myNotificationCategory"
	AutoCopy        = "0"
	IsArchive       = "1"
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
