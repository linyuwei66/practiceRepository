package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

type AppConfig struct{
	Release bool `ini:"release"`
	Port int `ini:"port"`
	*MySQLConfig `ini:"mysql"`
}

type MySQLConfig struct{
	User string `ini:"user"`
	Password string `ini:"password"`
	Host string `ini:"host"`
	Port int `ini:"port"`
	DB string `ini:"db"`
}

func Init(file string)error{
	return ini.MapTo(Conf,file)
}