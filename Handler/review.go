//创作影评的handler
package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/Func"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/model"
	"log"
	"sync"
	"time"
)

func Review(c *gin.Context) {
	log.Println("ReviewHandler start!")
	//解析token
	strToken := c.Param("token")
	claim, err := Func.VarifToken(strToken)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	var review model.Review
	if err := c.BindJSON(&review); err != nil {
		c.JSON(400, gin.H{
			"message": "Lost parameters",
		})
		return
	}

	//成功获取信息
	reviewID := strconv.Itoa(model.CountSum + 1)

	//获取userid的信息
	userInfo, err := Func.GetUserInfo(claim.UerID)

	//插入review表
	if err2 := Func.InsertReview(review, reviewID, userInfo, claim.UserID); err2 != nil {
		c.JSON(500, gin.H{
			"message": "server bussy",
		})
	}

	//请求成功，获取时间并返回
	t := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(200, gin.H{
		"review_id": reviewID,
		"time":      t,
	})
	model.CountSum++
}

func GetReview(c *gin.Context) {
	review_id := c.Param("review_id")
	if comment, err := GetReview(review_id); err != nil {
		c.JSON(500, gin.H{
			"message": "surver busy",
		})
		return
	}

	//成功获取，获取token
	strToken := c.Param("token")
	claim, err := Func.VarifToken(strToken)
	//如果没登陆，下面三个参数返回false
	if err != nil {
		c.JSON(200, gin.H{
			"comment":           comment,
			"comment_like":      false,
			"review_like":       false,
			"review_collection": false,
		})
        return
	}
	//登陆了还要获取参数

    com,rev,col := Func.GetExtraInfo(comment,claim.UserID,review_id)

	c.JSON(200, gin.H{
		"comment":           comment,
		"comment_like":      com,
		"review_like":       rev,
		"review_collection": col,
	})
}

func DeleteReview(c *gin.Context) {
    //先拿到id路径参数
    review_id := c.Param("review_id")
    //再查看token
    strToken := c.Param("token")
    claim, err := Func.VarifToken(strToken)
    if err != nil {
        c.JSON(401, gin.H{
            "message": "Wrong Token",
        })
        return
    }

    if err := DeleteFunc(claim.UserID,review_id);err != nil {
        c.JSON(400,gin.H{
            "message":"bad request",
        }
        return
    }

    c.JSON(200,gin.H{
        "message":"delete successfully",
    }
}

func ChangeReviewLike(c *gin.Context) {
    //先拿到id路径参数
    review_id := c.Param("review_id")
    //再查看token
    strToken := c.Param("token")
    claim, err := Func.VarifToken(strToken)
    if err != nil {
        c.JSON(401, gin.H{
            "message": "Wrong Token",
        })
        return
    }

    ChangeReviewLikeFunc(claim.UserID,review_id)
    c.JSON(200,gin.H{
        "message":"ok",
    })
}

func NewCollection(c *gin.Context) {
    //先拿到id路径参数
    review_id := c.Param("review_id")
    //再查看token
    strToken := c.Param("token")
    claim, err := Func.VarifToken(strToken)
    if err != nil {
        c.JSON(401, gin.H{
            "message": "Wrong Token",
        })
        return
    }

    NewCollection(claim.UserID,review_id)
    c.JSON(200,gin.H{
        "message":"ok",
    })
}
