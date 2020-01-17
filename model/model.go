package model

import(
	"strconv"
	"github.com/dgrijalva/jwt-go"
	"log"
	"errors"
)
type User struct{
	UserID    string  `json:"user_id"`
	UserPicture string `json:"user_picture"`
	Name  string      `json:"name"`
	Password  string  `json:"password"`
}

//是否存在该用户
func IfExistUser(id string) bool {
	var user = make([]User,1)
	if err:=Db.Self.Table("user").Where("user_id=?",id).First(&user).Error; err!=nil {
		log.Println(err.Error)
		return false 
	}
	if len(user)!=0{
		return true
	}
	return false
}


func Register( name string,password string,)string {
	count:=0
	Db.Self.Table("user").Count(&count)    //.Table() 是指定表名，不用默认表名
	temp:=202000+count
	id:=strconv.Itoa(temp)   //安全转换类型 直接强制转换不安全，可能会乱码

	user:=User{id,name,password}
	Db.Self.Table("user").Create(&user)
	return id
}

func VerifyPassword(id string, password string) bool{
	var user User
	Db.Self.Table("user").Where("user_id=?",id).First(&user)
	if user.Password==password {
		return true
	}
	log.Println(user.Password)
	log.Println(password)
	return false
}

//获取用户信息
type UserInfo struct{
	Followers   string `json:"followers"`
	Fans     	string  `json:"fans"`
	UserID     string   `json:"user_id"`
}

func GetFollowers( id string) string{
	count:=0 
	Db.Self.Table("follow").Where("user_id1=?",id).Count(&count)
	return strconv.Itoa(count)
}

func GetFans(id string) string {
	count:=0 
	Db.Self.Table("follow").Where("user_id2=?",id).Count(&count)	
	return strconv.Itoa(count)
}



// token生成和验证
type JwtClaims struct{
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

var	Secret = "sault"   //加盐

func GetToken(claims *JwtClaims) string {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
		return ""
	}	
	return signedToken
}

const(
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin = "请重新登陆"
)

func VerifyToken(strToken string) (*JwtClaims,error) {
//	strToken := c.Param("token")    //c.Param()可以获取单个参数,路径的也行

	//解析
	token, err := jwt.ParseWithClaims(strToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy+",或token解析失败")
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}

	return claims,nil  //JwtClaims结构体类型
}


