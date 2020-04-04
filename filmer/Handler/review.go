//创作影评的handler
package Handler

import (
	"github.com/filmer2/Func"
	"github.com/gin-gonic/gin"
	//"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
	"github.com/filmer2/model"
	"log"
	//"sync"
	"strconv"
	"time"
	//"net/http"
	//"github.com/dgrijalva/jwt-go"
	//"errors"
)

/*func Login(c *gin.Context) {
	claims := &model.JWTClaims{
		UserID:      "100001",
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(3600)).Unix()
	signedToken,err:=getToken(claims)
	if err!=nil{
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}*/

/*func getToken(claims *model.JWTClaims)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(model.Secret))
	if err != nil {
		return "",errors.New("ServerBusy")
	}
	return signedToken,nil
}*/

func Review(c *gin.Context) {
	log.Println("ReviewHandler start!")
	//解析token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	var review model.Review
	if err := c.BindJSON(&review); err != nil {
		c.JSON(400, gin.H{
			"message": "Lost parameters",
		})
		return
	}

	log.Println(review)
	//成功获取信息
	reviewID := model.Lastreviewid.Review_id + 1
	model.Lastreviewid.Review_id++

	//获取userid的信息
	userInfo, err := Func.GetUserInfo(claim.UserID)
	log.Println(userInfo)

	//插入review表
	if err2 := Func.InsertReview(review, reviewID, userInfo, claim.UserID); err2 != nil {
		c.JSON(500, gin.H{
			"message": "server bussy",
		})
		return
	}

	//请求成功，获取时间并返回
	t := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(200, gin.H{
		"review_id": reviewID,
		"time":      t,
	})
	model.Countsum.Countsum++
}

func GetReview(c *gin.Context) {
	log.Println("API START")
	review_id := c.Param("review_id")
	reviewID, _ := strconv.Atoi(review_id)
	log.Println(reviewID)
	var comment []model.CommentInfo
	var err error
	if comment, err = Func.GetReview(reviewID); err != nil {
		c.JSON(500, gin.H{
			"message": "surver busy",
		})
		return
	}

	//成功获取，获取token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
	//如果没登陆，下面三个参数返回false
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"comment":           comment,
			"review_like":       false,
			"review_collection": false,
		})
		return
	}
	//登陆了还要获取参数
	log.Println("have token")

	err,rev, col := Func.GetExtraInfo(&comment, claim.UserID, reviewID)
	if err != nil {
		c.JSON(500,gin.H{
			"message":"server busy",
		})
	}

	c.JSON(200, gin.H{
		"comment":           comment,
		"review_like":       rev,
		"review_collection": col,
	})
}

func DeleteReview(c *gin.Context) {
	//先拿到id路径参数
	review_id := c.Param("review_id")
	reviewID, _ := strconv.Atoi(review_id)
	//再查看token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	if err := Func.DeleteFunc(claim.UserID, reviewID); err != nil {
		c.JSON(400, gin.H{
			"message": "bad request",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "delete successfully",
	})
}

func ChangeReviewLike(c *gin.Context) {
	//先拿到id路径参数
	review_id := c.Param("review_id")
	reviewID, _ := strconv.Atoi(review_id)
	//再查看token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	if err2 := Func.ChangeReviewLikeFunc(claim.UserID, reviewID); err2 != nil {
		c.JSON(500, gin.H{
			"message": "server busy",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func NewCollection(c *gin.Context) {
	//先拿到id路径参数
	review_id := c.Param("review_id")
	reviewID, _ := strconv.Atoi(review_id)
	//再查看token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	if err2 := Func.NewCollection(claim.UserID, reviewID); err2 != nil {
		c.JSON(500, gin.H{
			"message": "ok",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func NewComment(c *gin.Context) {
	//先拿到id路径参数
	review_id := c.Param("review_id")
	reviewID, _ := strconv.Atoi(review_id)
	//解析token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{
			"message": "Lost parameters",
		})
		return
	}

	var commentID int
	var err1 error
	if err1, commentID = Func.NewComment(claim.UserID, reviewID, comment); err1 != nil {
		c.JSON(500, gin.H{
			"message": "surver busy",
		})
		return
	}

	//请求成功，获取时间并返回
	t := time.Now().Format("2006-01-02 15:04:05")
	c.JSON(200, gin.H{
		"comment_id": commentID,
		"time":       t,
	})
}

func NewCommentLike(c *gin.Context) {
	//先拿到id路径参数
	comment_id := c.Param("comment_id")
	commentID, _ := strconv.Atoi(comment_id)
	//再查看token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	claim, err := Func.VerifyToken(token.Token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	if err2 := Func.NewCommentLike(claim.UserID, commentID); err2 != nil {
		c.JSON(500, gin.H{
			"message": "server busy",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteComment(c *gin.Context) {
	//先拿到id路径参数
	comment_id := c.Param("comment_id")
	//再查看token
	var token model.StrToken
	c.BindHeader(&token)
	log.Println(token)
	_, err := Func.VerifyToken(token.Token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong Token",
		})
		return
	}

	commentID, err2 := strconv.Atoi(comment_id)
	if err2 != nil {
		c.JSON(500, gin.H{
			"message": "server busy",
		})
	}
	if err := Func.DeleteCommentFunc(commentID); err != nil {
		c.JSON(500, gin.H{
			"message": "server busy",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
