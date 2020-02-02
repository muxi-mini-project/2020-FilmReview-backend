package main

import (
	"github.com/gin-gonic/gin"
	"github.com/filmer/Func"
	"github.com/filmer/Handler"
	"github.com/filmer/database"
	"github.com/filmer/model"
	"log"
)

func init() {
	database.DB = database.InitDB(database.DB)
	model.Count = Func.CountInit(model.Count)
	model.Countsum = Func.CountSumInit(model.Countsum)
	model.Lastreviewid = Func.LastReviewIDInit(model.Lastreviewid)
	log.Println(model.Count,model.Countsum,model.Lastreviewid)
}

func main() {
	defer database.DB.Close()
	router := gin.Default()
	router.GET("/login", Handler.Login)
	router.POST("/review", Handler.Review)
	router.GET("/grounds", Handler.Grounds)
	router.GET("/grounds/:user_id", Handler.GroundsID)
	router.GET("/reviews/:review_id", Handler.GetReview)
	router.DELETE("/reviews/:review_id", Handler.DeleteReview)
	router.PUT("/reviews/:review_id", Handler.ChangeReviewLike)
	router.PATCH("/reviews/:review_id", Handler.NewCollection)
	router.POST("/reviews/:review_id/comment", Handler.NewComment)
	router.PUT("review/comments/:comment_id", Handler.NewCommentLike)
	router.DELETE("review/comments/:comment_id", Handler.DeleteComment)
	router.Run(":9091")
}
