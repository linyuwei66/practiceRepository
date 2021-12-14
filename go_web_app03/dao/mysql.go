package dao

import (
	"fmt"
	"go_web_app03/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	//astring := "root:123456@(127.0.0.1:3306)/goweb02?charset=utf8mb4&parseTime=True&loc=Local"
	astring := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",cfg.User,cfg.Password,cfg.Host,cfg.Port,cfg.DB)
	DB, err = gorm.Open("mysql", astring)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

