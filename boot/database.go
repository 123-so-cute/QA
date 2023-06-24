package boot

import (
	"QA_community/global"
	"QA_community/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitMysql() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		fmt.Println("连接mysql失败")
		return
	}
	global.GlobalDb = db
	global.GlobalDb.AutoMigrate(&model.User{}, &model.Question{}, &model.Answer{})

}

func InitRedis() {

}
