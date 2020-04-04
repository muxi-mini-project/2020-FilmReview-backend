package Handler

import (
	"github.com/filmer2/Func"
	"github.com/gin-gonic/gin"
	//"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
	"github.com/filmer2/model"
	"log"
)

func Grounds(c *gin.Context) {
	//收到请求直接返回结构体
	//一次性返回四个记录
	var groundInfos []model.GroundInfos
	var err error
	model.Count += 4                                                //4
	if groundInfos, err = Func.GetGround(model.Count); err != nil { //原来是model.Count
		c.JSON(500, gin.H{
			"message": "server busy",
		})
		return
	}

	//log.Println(groundInfos)
	//数据获取成功，现在返回
	c.JSON(200, groundInfos)
	log.Println(model.Count, model.Countsum.Countsum)

	if model.Count > model.Countsum.Countsum-model.Count { //model.Count>model.Countsum.Countsum-model.Count
		model.Count = -4 //-4
	}
}

func GroundsID(c *gin.Context) {
	//user_id := c.Param("user_id")
	//解析token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
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
			"message": "server busy",
		})
		return
	}

	c.JSON(200, ground)
}
