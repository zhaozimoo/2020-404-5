package config

import (

	"gopkg.in/gcfg.v1"
	"github.com/alecthomas/log4go"
)
//type MysqlConfig struct {
//	DbUserName string
//	DbPassword string
//	DbIp string
//	DbPort int
//	DbName string
//
//}

//type ServerInfo struct {
//	RCclienturl string
//	Reclienturl string
//}

type Return struct {
	DownloadUrl string
	DownServerPort int
	UploadServerPort string
}

type Config struct {
	//MysqlConfig MysqlConfig
	//ServerInfo  ServerInfo
	Return Return
}


func GetConfig() Config{
	var config Config
	err:=gcfg.ReadFileInto(&config,"Windmill/server/config.ini")
	if err != nil {
		log4go.Error(err)
	}
	return config
}

//func GetDBConfig() *mysql.DBConfig {
//	return mysql.NewMySqlConfig(GetConfig().MysqlConfig.DbUserName, GetConfig().MysqlConfig.DbPassword,GetConfig().MysqlConfig.DbIp,GetConfig().MysqlConfig.DbPort, GetConfig().MysqlConfig.DbName)
//}
