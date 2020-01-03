package config

import (
	"gopkg.in/gcfg.v1"
	"github.com/alecthomas/log4go"
)

type ServerInfo struct {
	User string
	Passwd string
	Host string
	Port int
	TargetIp string
	TargetPasswd string
	LocalPath string
	LocalDevPath string
	RemotePath string
	RemoteDevPath string
}

type Config struct {
	ServerInfo  ServerInfo
}

func GetConfig() Config{
	var config Config
	err:=gcfg.ReadFileInto(&config,"config.ini")
	if err != nil {
		log4go.Error(err)
	}
	return config
}