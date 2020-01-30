package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB
var Err error

func InitDB(DB *gorm.DB) *gorm.DB {
	log.Println("DB Connect Start")
	//连接数据库(考虑改成全局变量)
	DB, err := gorm.Open("mysql", "JacksieCheung:15811852133@/filmer?charset=utf8&parseTime=True&loc-Local")
	if err != nil {
		panic(err)
	}
	log.Println("DB Connect finished")
	return DB
}

func CloseDB() {
	DB.Close()
	return
}
