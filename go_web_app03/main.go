package main

// package go_web_app03

import (
	"fmt"
	"go_web_app03/dao"
	"go_web_app03/setting"

	"go_web_app03/models"

	"go_web_app03/routers"
)




func main() {

	setting.Init("conf/config.ini")

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	
	dao.DB.AutoMigrate(&models.Todo{})

	r:=routers.SetupRouter()
	// r.Run(":6060")
	r.Run(fmt.Sprintf(":%d",setting.Conf.Port))
}
