package main

import (
        _ "github.com/go-sql-driver/mysql"
        "github.com/jinzhu/gorm"
        "log"
//	"strconv"
)

type Review struct {
	Review_id int
}

func main() {
	var DB *gorm.DB
        DB, err := gorm.Open("mysql", "root:15811852133@/filmer?charset=utf8&parseTime=True&loc-Local")
        if err != nil {
                panic(err)
        }

        defer DB.Close()

	var review []Review
	var review2 []Review
	sql :="select review_id from user_review limit 1"
	DB.Raw(sql).Scan(&review)

	sql2 := "select review_id from user_review limit 2,1"
	DB.Raw(sql2).Scan(&review2)

	review = append(review,review2...)

	log.Println(review)
}
