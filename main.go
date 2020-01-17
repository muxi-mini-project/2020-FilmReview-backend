package main

import(
    "github.com/gin-gonic/gin"
    "github.com/ithub.com/muxi-mini-project/2020-FilmReview-backend/filmer/Handler"
    "github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
)

func init() {
    database.DB = database.InitDB(database.DB)
}

func main() {
    defer database.DB.Close()
    router:= gin.Default()
    //router.Post("路径",handler)
    router.Post("/review",Handler.Review)
    router.Get("/grounds",Handler.Grounds)
    router.Run("localhost:9090/api/v1")
}

