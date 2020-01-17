package model

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type Database struct{
	Self *gorm.DB
}

var Db *Database

func getDatabase() (*gorm.DB,error) {
	db,err :=gorm.Open("mysql", "root:120@/miniproject?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		log.Println(err)
	}
	return db,err
}

func (db *Database)Init(){
	newDb,_:=getDatabase()
	Db = &Database{Self:newDb}
} 

func (db *Database) Close(){
	Db.Self.Close()
}