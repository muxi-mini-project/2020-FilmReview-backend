package hander

import (
	"github.com/gin-gonic/gin"
	"github.com/project/model"
	"strconv"
	"time"
	//"log"
)

//　注册
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "输入有误，格式错误")
		return
	}
	user_id := model.Register(user.Name, user.Password)
	c.JSON(200, user_id)
}

//登陆
func Login(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "输入有误，格式错误")
		return
	}
	if !model.IfExistUser(user.UserID) {
		c.JSON(404, "用户名不存在")
		return
	}

	if !model.VerifyPassword(user.UserID, user.Password) {
		c.JSON(401, "密码错误")
		return
	}

	claims := &model.JwtClaims{UserID: user.UserID}
	//设置token过期时间
	ExpireTime := 3600000 // token有效期

	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	token := model.GetToken(claims)
	c.JSON(200, token)
}

//用户主页
func PeopleInfo(c *gin.Context) {
	id := c.Param("user_id") // c.Param()  是解析url里的参数
	token := c.Request.Header.Get("token")
	claims, err := model.VerifyToken(token)
	/*	if err!=nil{
			c.JSON(401,err.Error())
			return
		}
	*/ //发token　不验证
	followers := model.GetFollowers(id)
	fans := model.GetFans(id)
	attention := model.GetAttention(claims.UserID, id)
	user, err := model.GetUser(id)
	if err != nil {
		c.JSON(404, "找不到改用户信息")
		return
	}
	userInfo := model.UserInfo{Followers: followers, Fans: fans, UserID: user.UserID, UserPicture: user.UserPicture, Name: user.Name, Attention: attention}
	c.JSON(200, userInfo)
}

//修改用户信息
func UpdatePeopleInfo(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	if _, err := model.VerifyToken(token); err != nil {
		c.JSON(401, err.Error())
		return
	}
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, "输入有误，格式错误")
		return
	}
	user.UserID = id
	if err := model.UpdateUserInfo(user); err != nil {
		c.JSON(400, "更新失败")
		return
	}
	c.JSON(200, "修改成功")

}

//关注，取消关注用户
func Follow(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	claims, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}
	if err := model.Followone(claims.UserID, id); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "成功")
}

//我的影评
func MyReviews(c *gin.Context) {
	id := c.Param("user_id")
	reviews, err := model.GetReview(id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, reviews)
}

//创建专辑
func CreateAlbum(c *gin.Context) {
	var album model.Album
	token := c.Request.Header.Get("token")
	claims, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}
	if err := c.BindJSON(&album); err != nil {
		c.JSON(400, "输入有误，格式错误")
		return
	}

	album_id := model.NewAlbum(album, claims.UserID)
	c.JSON(200, album_id)
}

//移除专辑
func DeleteAlbums(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	var album_ids []model.Album
	if err := c.BindJSON(&album_ids); err != nil {
		c.JSON(400, "输入有误，格式错误")
		return
	}

	if err := model.DeleteAlbum(album_ids, claims.UserID); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "删除成功")
}

//用户主页－专辑
func Albums(c *gin.Context) {
	id := c.Param("user_id")
	albums, err := model.GetAlbums(id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, albums)
}

//专辑详情
func TheAlbum(c *gin.Context) {
	user_id := c.Param("user_id")
	album_id := c.Param("album_id")
	id, _ := strconv.Atoi(album_id)
	reviews, err := model.GetTheAlbum(user_id, id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, reviews)
}

//添加影评到专辑
func AddReviews(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	var album_review []model.AlbumReview
	if err := c.BindJSON(&album_review); err != nil {
		c.JSON(400, "输入有误，格式错误")
		return
	}
	if err := model.AddReviewsToAlbum(album_review); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "添加成功")

}

//从专辑中移除影评
func RemoveReviews(c *gin.Context) {
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, err.Error())
		return
	}

	var album_review []model.AlbumReview
	if err := c.BindJSON(&album_review); err != nil {
		c.JSON(400, "输入有误，格式错误")
		return
	}
	if err := model.RemoveReviewsFromAlbum(album_review); err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, "移除成功")

}

//用户主页－收藏
func Collection(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err1 := model.VerifyToken(token)
	if err1 != nil {
		c.JSON(401, err1.Error())
		return
	}

	reviews, err2 := model.GetCollection(claims.UserID)
	if err2 != nil {
		c.JSON(400, err2.Error())
		return
	}
	c.JSON(200, reviews)
}
