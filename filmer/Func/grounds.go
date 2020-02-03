//发现页面获取数据
package Func

import (
	"errors"
	//"github.com/gin-gonic/gin"
	"github.com/filmer/database"
	"github.com/filmer/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"sync"
)

func CountInit(count int) int {
	count = -4
	return count
}

func CountSumInit(countSum model.CountSum) model.CountSum {
	//countSum = 0
	sql := "select count(*) as countsum from user_review ;"
	database.DB.Debug().Raw(sql).Scan(&countSum)
	return countSum
}

func LastReviewIDInit(Lastreviewid model.LastReviewID) model.LastReviewID {
	sql := "select max(review_id) as review_id from user_review ;"
	database.DB.Debug().Raw(sql).Scan(&Lastreviewid)
	return Lastreviewid
}

//新版本
func GetGround(count int) ([]model.GroundInfos, error) {
	var User_review []model.GroundInfos
	sql := "select *from`user_review` order by review_id desc limit " + strconv.Itoa(count) + ",4;"
	log.Println(sql)
	if err := database.DB.Debug().Raw(sql).Scan(&User_review).Error; err != nil {
		return nil, errors.New("surver busy")
	}
	return User_review, nil
}

var lock sync.Mutex

//查看关注的界面
func GetGroundInfos(userid string, ground *[]model.GroundInfosID) {
	var ground1 []model.GroundInfosID
	log.Println("start")
	log.Println(userid)
	sql := "select name,user_picture,review_id,title,content,time,tag,picture,comment_sum,like_sum from user_review where user_id = " + userid
	log.Println(sql)
	database.DB.Raw(sql).Scan(&ground1)
	lock.Lock()
	(*ground) = append((*ground),ground1...)
	lock.Unlock()
	log.Println("finished")
	return
}

//获取关注的id
func GetUserID(userid string) ([]model.Follows, error) {
	var follows []model.Follows
	sql := "select user_id2 from follow where user_id1 = " + userid
	if err := database.DB.Raw(sql).Scan(&follows).Error; err != nil {
		log.Println(err)
		return follows, errors.New("surver busy")
	}
	return follows, nil
}

//我可以先获取关注的个数然后建立切片,最后合并。然后并发根据切片去找
func GetGroundAll(userid string) ([]model.GroundInfosID, error) {
	var follows []model.Follows
	var err error
	//先获取id
	if follows, err = GetUserID(userid); err != nil {
		return nil, errors.New("surver busy")
	}
	log.Println(follows)

	index := len(follows)
	log.Println(index)

	var ground []model.GroundInfosID
	if follows[0].User_id2 == "" {
		return ground,nil
	}


	wg := sync.WaitGroup{}
	wg.Add(index)
	var i int
	for i = 0; i < index; i++ {
		log.Println(follows[i])
		go func(i int) {
			GetGroundInfos(follows[i].User_id2, &ground)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return ground, nil
}
