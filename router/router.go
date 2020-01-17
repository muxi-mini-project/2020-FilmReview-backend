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
}



