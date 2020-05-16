package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

//var Err error

func InitDB(DB *gorm.DB) *gorm.DB {
	log.Println("DB Connect Start")
	//连接数据库(考虑改成全局变量)
	config := viper.New()
	config.AddConfigPath("$GOPATH/src/github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database") //手动改路径
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	if config.ReadInConfig(); err != nil {
		panic(err)
	}

	dns := fmt.Sprintf("%s:%s@/filmer?charset=utf8&parseTime=True&loc-Local",
		config.Get("database.user"),
		config.Get("database.pwd"))

	DB, err := gorm.Open("mysql", dns)
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
