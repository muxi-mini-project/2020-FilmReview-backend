package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/Handler"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
)

func init() {
	database.DB, database.Err = database.InitDB(database.DB)
	if database.Err != nil {
		panic(database.Err)
	}
	Func.CountInit(model.Count)
	Func.CountSumInit(model.CountSum)
}

func main() {
	defer database.DB.Close()
	router := gin.Default()
	//router.Post("路径",handler)
	router.Post("/review", Handler.Review)
	router.Get("/grounds", Handler.Grounds)
	router.Get("/grounds/:user_id", Handler.GroundsID)
	router.Get("/reviews/:review_id", Handler.GetReview)
	router.Delete("/reviews/:review_id", Handler.DeleteReview)
	router.Put("/reviews/:review_id", Handler.ChangeReviewLike)
	router.Patch("/reviews/:review_id", Handler.NewCollection)
	router.Post("/reviews/:review_id/comment", Handler.NewComment)
	router.Put("/reviews/comments/:comment_id", Handler.NewCommentLike)
	router.Delete("/reviews/comments/:comment_id", Hadnler.DeleteComment)
	router.Run("localhost:9090/api/v1")
}
