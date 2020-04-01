package SearchApi

import (
	"github.com/filmer/model"
	"github.com/gin-gonic/gin"
	"log"
)

func Searcher(c *gin.Context) {
	var sentence word
	var words []string
	var result []model.GroundInfosID

	if err := c.BindJSON(&sentence); err != nil {
		c.JSON(401, gin.H{
			"message": "try it again in the right way",
		})
		return
	}
	words = cutWords([]byte(sentence.Words))

	log.Println(words)
	//并发错误处理！
	err := doSearch(words, &result)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
	}

	c.JSON(200,result)
}

func ReturnSearch(c *gin.Context) {
	var sentence word
	var words []string
	var resultTags []tag

	if err := c.BindJSON(&sentence); err != nil {
		c.JSON(401, gin.H{
			"message": "try it again in the right way",
		})
		return
	}
	words = cutWords([]byte(sentence.Words))

	log.Println(words)
	//并发错误处理！
	err := doReturnSearch(words, &resultTags)
	if err != nil || len(resultTags) == 0 {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.JSON(200,resultTags)
}
