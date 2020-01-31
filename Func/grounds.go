//发现页面获取数据
package Func

import (
	"errors"
	//"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/model"
	//"log"
	"strconv"
	"sync"
)

func CountInit(count int) {
	count = -3
	return
}

func CountSumInit(countSum int) {
	countSum = 100000
	return
}

//一次返回四张
//这是正常情况
/*func GetGround(count int) ([]model.GroundInfos,error) {
    var groundInfos []model.GroundInfos
    sql:= "select review_id，user_id from myreviews limit "+strconv.Iota(count)+",4"
    if err := DB.Raw(sql).Scan(&groundInfos).Error();err != nil {
        return errors.New("survery bussy")
    }
    return myReview
}
//查询同时编入结构体里
func GetUser(userID string,groundInfos []model.GroundInfos,index int) error {
    sql:= "select user_picture,name from user where user_id = "+userID

    if err := database.DB.Raw(sql).Scan(&groundInfos[index]).Error();err != nil {
        return errors.New("server bussy")
    }
    return nil
}

func GetReview(reviewID string,groundInfos []model.GroundInfos,index int) error {
    sql := "select title,content,time,tag,picture,like_sum from review where reviewID = "+reviewID
    if err := database.DB.Raw(sql).Scan(&groundInfos[index]).Error();err != nil {
        return errors.New("server bussy")
    }
    return nil
}

func GetComment(reviewID string,groundInfos []model.GroundInfos,index int) error {
    sql := "select count(*) from user_com where review_id = "+reviewID
    if err := database.DB.Raw(sql).Scan(&groundInfos[index].Error(); err != nil {
        return errors.New("server bussy")
    }
    return nil
}

func GetGroundInfos(count int) ([]model.GroundInfos,err) {
    if groundInfos,err := GetGround(count);err != nil {
        return errors.New("survery bussy")
    }
    //拿到了要返回的reviewid和userid，接下来要并发爬user表和review表
    index := len(groundInfos)

    wg := sync.WaitGroup{}
    wg.Add(index*3)

    var l sync.Mutex
    for i := 0; i < index ; i++ {
        go func() {
            l.Lock()
            if err := GetUser(groundInfos[i].review_id,groundInfos,i);err != nil {
                return errors.New("surver bussy")
            }
            l.Unlock()
            wg.Done()
        }()

        go func() {
            l.Lock()
            if err := GetReview(groundInfos[i].review_id,groundInfos,i);err != nil {
                return errors.New("surver bussy")
            }
            l.Unlock()
            wg.Done()
        }()

        go func() {
            l.Lock()
            if err := GetComment(groundInfos[i].review_id,groundInfos,i);err != nil {
                return errors.New("surver bussy")
            }
            l.Unlock()
            wg.Done()
        }()
    }

    wg.Wait()

    return groundInfos,nil
}

func GetGroundID() ([]model.GroundInfos,error) {
    var groundInfos []model.GroundInfos
    sql := "select user_id from "*/

//新版本
func GetGround(count int) ([]model.GroundInfos, error) {
	var ground []model.GroundInfos
	sql := "select review_id,title,content,time,tag,picture,comment_sum,like_sum from user_review limit " + strconv.Itoa(count) + ",4"
	if err := database.DB.Raw(sql).Scan(&ground).Error; err != nil {
		return nil, errors.New("surver busy")
	}
	return ground, nil
}

//查看关注的界面
func GetGroundInfos(userid string, ground model.GroundInfosID) error {
	sql := "select name,user_picture,review_id,title string,content,time,tag,picture,comment_sum,like_sum from user_review where user_id = " + userid
	if err := database.DB.Raw(sql).Scan(&ground).Error; err != nil {
		return errors.New("surver busy")
	}
	return nil
}

//获取关注的id
func GetUserID(userid string) ([]model.GroundInfosID, error) {
	var ground []model.GroundInfosID
	sql := "select user_id2 from follow where user_id1 = " + userid
	if err := database.DB.Raw(sql).Pluck("use_id2", &ground).Error; err != nil {
		return nil, errors.New("surver busy")
	}
	return ground, nil
}

func GetGroundAll(userid string) ([]model.GroundInfosID, error) {
	var ground []model.GroundInfosID
	var err error
	//先获取id
	if ground, err = GetUserID(userid); err != nil {
		return nil, errors.New("surver busy")
	}

	//接下来通过并发获取每个id的review
	index := len(userid)
	wg := sync.WaitGroup{}
	wg.Add(index)

	for i := 0; i < index; i++ {
		go func() {
			if err2 := GetGroundInfos(ground[i].User_id2, ground[i]); err2 != nil {
				//return nil, errors.New("surver bussy")//并发错误处理
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return ground, nil
}
