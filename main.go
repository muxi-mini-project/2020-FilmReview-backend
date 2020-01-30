package main

import (
	"fmt"
	"github.com/project/model"
	"github.com/project/router"
)

func main() {
	model.Db.Init()
	defer model.Db.Close()

	router.InitRouter()
	fmt.Println("这是一条友爱的中文版二进制分割线,老祖宗保佑别出bug，别出bug,别出bug,别出bug")
	fmt.Println("阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳阴阳")
	router.Router.Run(":8080")
}
