package SearchApi

import (
    "github.com/gin-gonic/gin"
    "github.com/filmer/model"
    //"log"
)

func Searcher(c *gin.Context) {
    var sentence word
    var words []string
    var result []model.GroundInfosID

    words = cutWords([]byte(sentence.Word))

    //并发错误处理！
    err := doSearch(words,&result)
    if err != nil {
        c.JSON(404,gin.H{
            "message":"Not Found",
        })
    }

    c.JSON(200,gin.H{
        "result":result,
    })
}

func ReturnSearch(c *gin.Context) {
    var sentence word
    var words []string
    var resultTags []tag

    words = cutWords([]byte(sentence.Word))

    //并发错误处理！
    err := doReturnSearch(words,&resultTags)
    if err != nil {
        c.JSON(404,gin.H{
            "message":"Not Found",
        })
    }

    c.JSON(200,gin.H{
        "resultTags":resultTags,
    })
}
