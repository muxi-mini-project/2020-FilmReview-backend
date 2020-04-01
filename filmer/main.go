package main

import (
	"github.com/filmer/Func"
	"github.com/filmer/Handler"
	"github.com/filmer/database"
	"github.com/filmer/model"
	"github.com/gin-gonic/gin"
	//"github.com/filmer/modelWency"
	"github.com/filmer/SearchApi"
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

	router.POST("/api/v1/createuser", Handler.CreateUser)
	router.POST("/api/v1/login", Handler.Login)

	g1 := router.Group("/api/v1/grounds")
	{
		g1.GET("/", Handler.Grounds)
		g1.GET("/:user_id", Handler.GroundsID)
	}

	g2 := router.Group("/api/v1/reviews")
	{
		g2.GET("/:review_id", Handler.GetReview)
		g2.DELETE("/:review_id", Handler.DeleteReview)
		g2.PUT("/:review_id", Handler.ChangeReviewLike)
		g2.PATCH("/:review_id", Handler.NewCollection)
		g2.POST("/:review_id/comment", Handler.NewComment)
	}

	g5 := router.Group("api/v1/review")
	{
		g5.PUT("/comments/:comment_id", Handler.NewCommentLike)
		g5.DELETE("/comments/:comment_id", Handler.DeleteComment)
		g5.POST("/", Handler.Review)
	}

	g3 := router.Group("/api/v1/searcher")
	{
		g3.POST("/results", SearchApi.Searcher)
		g3.POST("/tags", SearchApi.ReturnSearch)
	}

	g4 := router.Group("/api/v1/people")
	{
		g4.GET("/:user_id", Handler.PeopleInfo)
		g4.PUT("/:user_id", Handler.UpdatePeopleInfo)
		g4.PATCH("/:user_id", Handler.Follow)
		g4.GET("/:user_id/myreviews", Handler.MyReviews)
		g4.POST("/:user_id/albums", Handler.CreateAlbum)
		g4.GET("/:user_id/albums", Handler.Albums)
		g4.PUT("/:user_id/albums", Handler.DeleteAlbums)
		g4.GET("/:user_id/albums/:album_id", Handler.TheAlbum)
		g4.POST("/:user_id/albums/:album_id", Handler.AddReviews)
		g4.PUT("/:user_id/albums/:album_id", Handler.RemoveReviews)
		g4.GET("/:user_id/collection", Handler.Collection)
	}

	router.Run(":9092")
}
