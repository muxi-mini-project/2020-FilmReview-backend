package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/Func"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/Handler"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/model"
)

func init() {
	database.DB = database.InitDB(database.DB)
	Func.CountInit(model.Count)
	Func.CountSumInit(model.CountSum)
}

func main() {
	defer database.DB.Close()
	router := gin.Default()
	//router.Post("路径",handler)
	router.POST("/review", Handler.Review)
	router.GET("/grounds", Handler.Grounds)
	router.GET("/grounds/:user_id", Handler.GroundsID)
	router.GET("/reviews/:review_id", Handler.GetReview)
	router.DELETE("/reviews/:review_id", Handler.DeleteReview)
	router.PUT("/reviews/:review_id", Handler.ChangeReviewLike)
	router.PATCH("/reviews/:review_id", Handler.NewCollection)
	router.POST("/reviews/:review_id/comment", Handler.NewComment)
	router.PUT("/reviews/comments/:comment_id", Handler.NewCommentLike)
	router.DELETE("/reviews/comments/:comment_id", Handler.DeleteComment)
	router.Run("localhost:9090/api/v1")
}
