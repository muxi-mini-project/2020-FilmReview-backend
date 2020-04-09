package Handler

import (
	"github.com/filmer2/modelWency"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

//　注册
func CreateUser(c *gin.Context) {
	var user modelWency.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误",
		})
		return
	}
	userId := modelWency.Register(user.Name, user.Password)
	//log.Println(userId)
	c.JSON(200, gin.H{
		"userId": userId,
	})
}

//登陆
func Login(c *gin.Context) {
	var user modelWency.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	if !modelWency.IfExistUser(user.UserID) {
		c.JSON(404, gin.H{"message": "用户名不存在"})
		return
	}

	if !modelWency.VerifyPassword(user.UserID, user.Password) {
		c.JSON(401, gin.H{"message": "密码错误"})
		return
	}

	claims := &modelWency.JwtClaims{UserID: user.UserID}
	log.Println(claims)
	//设置token过期时间
	ExpireTime := 3600000 // token有效期

	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	token := modelWency.GetToken(claims)
	c.JSON(200, gin.H{"token": token})
}

//用户主页
func PeopleInfo(c *gin.Context) {
	id := c.Param("user_id") // c.Param()  是解析url里的参数
	token := c.Request.Header.Get("token")
	claims, err := modelWency.VerifyToken(token)
	/*	if err!=nil{
			c.JSON(401,err.Error())
			return
		}
	*/ //发token　不验证
	followers := modelWency.GetFollowers(id)
	fans := modelWency.GetFans(id)
	attention := modelWency.GetAttention(claims.UserID, id)
	user, err := modelWency.GetUser(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到改用户信息"})
		return
	}
	userInfo := modelWency.UserInfo{Followers: followers, Fans: fans, UserID: user.UserID, UserPicture: user.UserPicture, Name: user.Name, Attention: attention}
	c.JSON(200, userInfo)
}

//修改用户信息
func UpdatePeopleInfo(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	if _, err := modelWency.VerifyToken(token); err != nil {
		c.JSON(401, err.Error())
		return
	}
	var user modelWency.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	user.UserID = id
	if err := modelWency.UpdateUserInfo(user); err != nil {
		c.JSON(400, gin.H{"message": "更新失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})

}

//关注，取消关注用户
func Follow(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	claims, err := modelWency.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}
	if err := modelWency.Followone(claims.UserID, id); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "成功"})
}

//我的影评
func MyReviews(c *gin.Context) {
	id := c.Param("user_id")
	reviews, err := modelWency.GetReview(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200,  reviews)
}

//创建专辑
func CreateAlbum(c *gin.Context) {
	var album modelWency.Album
	token := c.Request.Header.Get("token")
	claims, err := modelWency.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}
	if err := c.BindJSON(&album); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}

	albumId := modelWency.NewAlbum(album, claims.UserID)
	c.JSON(200, gin.H{"albumId": albumId})
}

//移除专辑
func DeleteAlbums(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err := modelWency.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	var albumIds []modelWency.Album
	if err := c.BindJSON(&albumIds); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}

	if err := modelWency.DeleteAlbum(albumIds, claims.UserID); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

//用户主页－专辑
func Albums(c *gin.Context) {
	id := c.Param("user_id")
	albums, err := modelWency.GetAlbums(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, albums)
}

//专辑详情
func TheAlbum(c *gin.Context) {
	userId := c.Param("userId")
	albumId := c.Param("albumId")
	id, _ := strconv.Atoi(albumId)
	reviews, err := modelWency.GetTheAlbum(userId, id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, reviews)
}

//添加影评到专辑
func AddReviews(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := modelWency.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	var albumReview []modelWency.AlbumReview
	if err := c.BindJSON(&albumReview); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	if err := modelWency.AddReviewsToAlbum(albumReview); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "添加成功"})

}

//从专辑中移除影评
func RemoveReviews(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := modelWency.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": err.Error()})
		return
	}

	var albumReview []modelWency.AlbumReview
	if err := c.BindJSON(&albumReview); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	if err := modelWency.RemoveReviewsFromAlbum(albumReview); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{"message": "移除成功"})

}

//用户主页－收藏
func Collection(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err1 := modelWency.VerifyToken(token)
	if err1 != nil {
		c.JSON(401, gin.H{"message": err1.Error()})
		return
	}

	reviews, err2 := modelWency.GetCollection(claims.UserID)
	if err2 != nil {
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	}
	c.JSON(200, reviews)
}
