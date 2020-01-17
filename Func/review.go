//存放创建review的函数
package Func

import (
    "strconv"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
    "github.com/muxi-mini-project/2020-FilmReview-backend/filmer/model"
    _ "github.com/go-sql-driver/mysql"
)

//通过全局变量CountSum实现
/*func CountRecord() (string,error) {
    var count Count
    sql := "select count(*) from review"
    err:= database.DB.Raw(sql).Scan(&count).Error()

    count.count=strconv.Atoi(count.count)
    return stronv.Itoa(100000+count.Count),err
}*/

func InsertReview(review model.Review,reviewID string) error {
    sql := "insert into review(review_id,title,content,time,tag,picture,like_sum) values("+reviewID+","+review.title+","+review.context+","+review.tag+","+review.picture+","+review.like_sum+")"
    err := database.DB.Raw(sql).Error()
    return err
}

func InsertMyReview(reviewID,userID string) error {
    sql := "insert into myreviews(user_id,review_id) values("+reviewID+","+userID+")"
    err := database.DB.Raw(sql).Error()
    return err
}

