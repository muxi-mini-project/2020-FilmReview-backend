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

	wg := sync.WaitGroup{}
	wg.Add(2)

	//插入review表
	go func() {
		if err2 := Func.InsertReview(review, reviewID); err2 != nil {
			c.JSON(500, gin.H{
				"message": "server bussy",
			})
		}
		wg.Done()
	}()
	//插入review和user的关系表
	go func() {
		if err3 := Func.InsertMyReview(reviewID, claim.UserID); err3 != nil {
			c.JSON(500, gin.H{
				"message": "server bussy",
			})
		}
		wg.Done()
	}()

	wg.Wait()

	//请求成功，获取时间并返回
	t := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(200, gin.H{
		"review_id": reviewID,
		"time":      t,
	})
	model.CountSum++
}
