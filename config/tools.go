package config

func GetDsn() string {
	return LocalConfig.Mysql.Host + ":" +
		LocalConfig.Mysql.Port +
		"@tcp(" +
		LocalConfig.Mysql.Host + ":" +
		LocalConfig.Mysql.Port + ")/" +
		LocalConfig.System.Name +
		"?charset=utf8mb4&parseTime=True&loc=Local"

}
