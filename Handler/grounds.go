package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/Func"
	//"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/model"
	//"log"
)

func Grounds(c *gin.Context) {
	//收到请求直接返回结构体
	//一次性返回四个记录
	var groundInfos []model.GroundInfos
	var err error
	model.Count += 4
	if groundInfos, err = Func.GetGround(model.Count); err != nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}

	//数据获取成功，现在返回
	c.JSON(200, gin.H{
		"Found": groundInfos,
	})

	if model.Count > model.CountSum {
		model.Count = 0
	}
}

func GroundsID(c *gin.Context) {
	//user_id := c.Param("user_id")
	//解析token
	strToken := c.Param("token")
	claim, err := Func.VerifyToken(strToken)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	var ground []model.GroundInfosID
	var err2 error
	if ground, err2 = Func.GetGroundAll(claim.UserID); err2 != nil {
		c.JSON(500, gin.H{
			"message": "surver busy",
		})
		return
	}

	c.JSON(200, gin.H{
		"Ground": ground,
	})
}
