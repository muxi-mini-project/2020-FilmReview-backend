package main

import (
	"github.com/filmer/Func"
	"github.com/filmer/Handler"
	"github.com/filmer/database"
	"github.com/filmer/model"
	"github.com/gin-gonic/gin"
	//"github.com/filmer/modelWency"
	"log"
)

func init() {
	database.DB = database.InitDB(database.DB)
	model.Count = Func.CountInit(model.Count)
	model.Countsum = Func.CountSumInit(model.Countsum)
	model.Lastreviewid = Func.LastReviewIDInit(model.Lastreviewid)
	log.Println(model.Count, model.Countsum, model.Lastreviewid)
}

func main() {
	defer database.DB.Close()
	router := gin.Default()
	router.POST("/api/v1/review", Handler.Review)
	router.GET("/api/v1/grounds", Handler.Grounds)
	router.GET("/api/v1/grounds/:user_id", Handler.GroundsID)
	router.GET("/api/v1/reviews/:review_id", Handler.GetReview)
	router.DELETE("/api/v1/reviews/:review_id", Handler.DeleteReview)
	router.PUT("/api/v1/reviews/:review_id", Handler.ChangeReviewLike)
	router.PATCH("/api/v1/reviews/:review_id", Handler.NewCollection)
	router.POST("/api/v1/reviews/:review_id/comment", Handler.NewComment)
	router.PUT("/api/v1/review/comments/:comment_id", Handler.NewCommentLike)
	router.DELETE("/api/v1/review/comments/:comment_id", Handler.DeleteComment)

	router.POST("/api/v1/createuser", Handler.CreateUser)
	router.POST("/api/v1/login", Handler.Login)
	router.GET("/api/v1/people/:user_id", Handler.PeopleInfo)
	router.PUT("/api/v1/people/:user_id", Handler.UpdatePeopleInfo)
	router.PATCH("/api/v1/people/:user_id", Handler.Follow)
	router.GET("/api/v1/people/:user_id/myreviews", Handler.MyReviews)
	router.POST("/api/v1/people/:user_id/albums", Handler.CreateAlbum)
	router.GET("/api/v1/people/:user_id/albums", Handler.Albums)
	router.PUT("/api/v1/people/:user_id/albums", Handler.DeleteAlbums)
	router.GET("/api/v1/people/:user_id/albums/:album_id", Handler.TheAlbum)
	router.POST("/api/v1/people/:user_id/albums/:album_id", Handler.AddReviews)
	router.PUT("/api/v1/people/:user_id/albums/:album_id", Handler.RemoveReviews)
	router.GET("/api/v1/people/:user_id/collection", Handler.Collection)
	router.Run(":9091")
}
