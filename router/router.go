package router

import (
	"github.com/gin-gonic/gin"
	"github.com/project/hander"
)

var Router *gin.Engine

func InitRouter(){
	Router= gin.Default()
	Router.POST("/api/v1/createuser",hander.CreateUser)
	Router.POST("/api/v1/login",hander.Login)
	Router.GET("/api/v1/people/:user_id",hander.PeopleInfo)
	Router.PUT("/api/v1/people/:user_id",hander.UpdatePeopleInfo)
	Router.PATCH("/api/v1/people/:user_id",hander.Follow)
	Router.GET("/api/v1/people/:user_id/myreviews",hander.MyReviews)
	Router.POST("/api/v1/people/:user_id/albums",hander.CreateAlbum)
	Router.GET("/api/v1/people/:user_id/albums",hander.Albums)
	Router.PUT("/api/v1/people/:user_id/albums",hander.DeleteAlbums)
	Router.GET("/api/v1/people/:user_id/albums/:album_id",hander.TheAlbum)
	Router.POST("/api/v1/people/:user_id/albums/:album_id",hander.AddReviews)
	Router.PUT("/api/v1/people/:user_id/albums/:album_id",hander.RemoveReviews)
}



