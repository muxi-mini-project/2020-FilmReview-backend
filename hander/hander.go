package hander

import (
	"github.com/gin-gonic/gin"
	"github.com/project/model"
	"time"
	//"log"
)


//　注册
func CreateUser(c *gin.Context){
	var user model.User
	if err:=c.BindJSON(&user); err !=nil{
		c.JSON(400,"输入有误，格式错误")
		return
	}
	user_id:=model.Register(user.Name,user.Password)
	c.JSON(200,user_id)
}


//登陆
func Login(c *gin.Context){
	var user model.User
	if err:=c.BindJSON(&user); err !=nil{
		c.JSON(400,"输入有误，格式错误")
		return
	}
	if !model.IfExistUser(user.UserID){
		c.JSON(404,"用户名不存在")
		return
	}

	if !model.VerifyPassword(user.UserID,user.Password) {
		c.JSON(401,"密码错误")
		return
	}

	claims:= &model.JwtClaims{UserID:user.UserID}
	//设置token过期时间
	ExpireTime:=3600    // token有效期

	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	token:=model.GetToken(claims)
	c.JSON(200,token)
}


//用户主页
func PeopleInfo(c *gin.Context){
	 //token:=c.Param("token")  // c.Param()  是解析url里的参数
	token:=c.Request.Header.Get("token")
	claims,err:=model.VerifyToken(token)
	if err!=nil{
		c.JSON(401,err.Error())
		return
	}
	
	followers:=model.GetFollowers(claims.UserID)
	fans:=model.GetFans(claims.UserID)
	userInfo:=model.UserInfo{Followers:followers,Fans:fans,UserID:claims.UserID}
	c.JSON(200,userInfo)
}


//修改用户信息
func UpdateInfo(){
	token:=c.Request.Header.Get("token")
	claims,err:=model.VerifyToken(token)
	if err!=nil{
		c.JSON(401,err.Error())
		return
	}
	var user model.User
	if err:=c.BindJSON(&user); err !=nil{
		c.JSON(400,"输入有误，格式错误")
		return
	}


}


