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
		log.Println(err.Error())
		return false 
	}
	if len(user)!=0{
		return true
	}
	return false
}


//注册
func Register( name string,password string,)string {
	count:=0
	if err:=Db.Self.Table("user").Count(&count).Error; err!=nil{    //.Table() 是指定表名，不用默认表名
		log.Println("RegisterError"+err.Error())
		return ""
	}
	//这种情况不能删除记录，不然ID会重复
	temp:=202000+count
	id:=strconv.Itoa(temp)   //安全转换类型 直接强制转换不安全，可能会乱码

	user:=User{UserID:id,Name:name,Password:password}
	if err:= Db.Self.Table("user").Create(&user).Error; err!=nil{
		log.Println("RegisterError"+err.Error())
		return ""
	}
	return id
}


//登录验证
func VerifyPassword(id string, password string) bool{
	var user User
	if err:=Db.Self.Table("user").Where("user_id=?",id).First(&user).Error; err!=nil{
		log.Println("VerifyPasswordError"+err.Error())
		return false
	}
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
	UserID    string  `json:"user_id"`
	UserPicture string `json:"user_picture"`
	Name  string      `json:"name"`
	Attention  bool   `json:"attention"`
}

func GetFollowers( id string) string{
	count:=0 
	if err:=Db.Self.Table("follow").Where("user_id1=?",id).Count(&count).Error; err!=nil{
		log.Println("GetFollowersError"+err.Error())
		return ""
	}
	return strconv.Itoa(count)
}

func GetFans(id string) string {
	count:=0 
	if err:=Db.Self.Table("follow").Where("user_id2=?",id).Count(&count).Error; err!=nil{
		log.Println("GetFansError"+err.Error())
		return ""
	}
	return strconv.Itoa(count)
}

func GetAttention(id1 string, id2 string) bool {
	count:=0 
	if err:=Db.Self.Table("follow").Where(map[string]interface{}{"user_id1":id1, "user_id2": id2}).Count(&count).Error; err!=nil{
		log.Println("GetFansError"+err.Error())
		return false
	}
	if count==1 {
		return true
	}
	return false
}

func GetUser(id string) (User,error) {
	var user User
	if err:=Db.Self.Table("user").Where("user_id=?",id).First(&user).Error; err!=nil{
		log.Println("GetUser"+err.Error())
		return User{},err
	}
	return user,nil
}


//修改用户信息
func UpdateUserInfo(user User) error {
	// 使用 `map` 更新多个属性，只会更新那些被更改了的字段
	//log.Println(user)
	if err:=Db.Self.Table("user").Where("user_id=?",user.UserID).Updates(map[string]interface{}{"name": user.Name, "user_picture": user.UserPicture}).Error; err!=nil{
		log.Println("UpdateUserInfo"+err.Error())
		return err
	}
	return nil
}

type Follower struct{
	UserID1   string
	UserID2   string
}

//关注，取消关注用户
func Followone(id1 string, id2 string) error {
	fans:=Follower{UserID1:id1,UserID2:id2}
	if GetAttention(id1,id2){
		if err:=Db.Self.Table("follow").Delete(&fans).Error;err!=nil{
			log.Println("Followone"+err.Error())
			return err
		}
		return nil
	}

	if err:= Db.Self.Table("follow").Create(&fans).Error; err!=nil{
		log.Println("Followone"+err.Error())
		return err
	}
	return nil

}



type Review struct{
	UserID   string      `json:"user_id"`
	Name 	 string      `json:"name"`
	UserPicture  string  `json:"user_picture"`
	ReviewId    string   `json:"review_id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Time		string   `json:"time"`
	Tag 		string   `json:"tag"`
	Picture  	string   `json:"picture"`
	CommentSum  int   `json:"comment_sum"`
	LikeSum 	int   `json:"like_sum"`
}

//用户主页的我的影评
func GetReview(id string) ( []Review,error) {
	var reviews []Review
	if err:=Db.Self.Table("user_review").Where("user_id=?",id).Find(&reviews).Error;err!=nil{
		log.Println("GetReview"+err.Error())
		return nil,err
	}
	return reviews,nil
}


type Album struct{
	UserID  string  `json:"user_id"`
	AlbumID string  `json:"album_id"`
	Title   string  `json:"title"`
	Summary string  `json:"summary"`
	ContentSum int  `json:"content_sum"`
}

//创建新专辑
func NewAlbum(album Album, id string) string {
	var lastalbum Album
	if err:=Db.Self.Table("album").Where("user_id=?",id).Order("album_id desc").First(&lastalbum).Error; err!=nil{   
		//没找到记录也会报错
		lastalbum.AlbumID="0"
}
	album.UserID=id
	temp,_:=strconv.Atoi(lastalbum.AlbumID)
	album.AlbumID=strconv.Itoa(temp+1)
	album.ContentSum=0
	if err:= Db.Self.Table("album").Create(&album).Error; err!=nil{
		log.Println("NewAlbum "+err.Error())
		return "出错"
	}
	return album.AlbumID
}

//用户主页－专辑
func GetAlbums(id string) ([]Album, error){
	var albums []Album
	if err:=Db.Self.Table("album").Where("user_id=?",id).Find(&albums).Error;err!=nil{
		log.Println("GetAlbums"+err.Error())
		return nil,err
	}
	return albums,nil
}


//移除专辑
func DeleteAlbum(album_ids []Album, user_id string) error {
	for i:=0; i<len(album_ids); i++{                                                           //注意这里的album_ids,用其他的报错了,类型可以是数组或者结构体
		if err:=Db.Self.Table("album").Where("user_id=? AND album_id=?",user_id,album_ids[i].AlbumID).Delete(&album_ids).Error;err!=nil{
			log.Println("DeleteAlbum"+err.Error())
			return err
		}		  //没找到不会报错
	}
	return nil
}

//专辑详情
func GetTheAlbum(user_id string, album_id string) ([]Review, error) {
	var review_ids []string
	if err:=Db.Self.Table("album_review").Where("user_id=? AND album_id=?",user_id,album_id).Pluck("review_id",&review_ids).Error;err!=nil{
		log.Println("GetTheAlbum1 "+err.Error())
		return nil,err
	}
//设定了切片初始长度,和容量,直接声明切片会index超出
	reviews:=make([]Review,len(review_ids),30)
	for i:=0; i<len(review_ids); i++{                        
		if err:=Db.Self.Table("user_review").Where("review_id=?",review_ids[i]).Find(&reviews[i]).Error;err!=nil{
			log.Println("GetTheAlbum2 "+err.Error())
			return nil,err
		}
	}
	return reviews,nil
}


type AlbumReview struct{
	UserID    string `json:"user_id"`   
	AlbumID   string  `json:"album_id"`
	ReviewID  string  `json:"review_id"`
}

//添加影评到专辑
func AddReviewsToAlbum(album_review []AlbumReview) error {
	//                                   注意这个直接是数组切片
	for i:=0; i<len(album_review); i++{
		if err:= Db.Self.Table("album_review").Create(&album_review[i]).Error; err!=nil{
			log.Println("AddReviewsToAlbum1 "+err.Error())
			return err
		}
	}

	//专辑中的contentsum要更新
	var album  Album
	if err:=Db.Self.Table("album").Where("user_id=? AND album_id=?",album_review[0].UserID,album_review[0].AlbumID).Find(&album).Error;err!=nil{
		log.Println("AddReviewsToAlbum2 "+err.Error())
		return err
	}

	if err:=Db.Self.Table("album").Where("user_id=? AND album_id=?",album.UserID,album.AlbumID).Update("content_sum",album.ContentSum+len(album_review)).Error; err!=nil{
		log.Println("AddReviewsToAlbum3 "+err.Error())
		return err
	}
	return nil
}

//从专辑中移除影评
func RemoveReviewsFromAlbum(album_review []AlbumReview) error{
	for i:=0; i<len(album_review); i++{
											//可以连用AND
		if err:=Db.Self.Table("album_review").Where("user_id=? AND album_id=? And review_id=?",album_review[0].UserID,album_review[0].AlbumID,album_review[i].ReviewID).Delete(&album_review).Error;err!=nil{
			log.Println("RemoveReviewsFromAlbum1 "+err.Error())
			return err
		}		  //没找到不会报错
	}
	//专辑中的contentsum要更新
	var album  Album
	if err:=Db.Self.Table("album").Where("user_id=? AND album_id=?",album_review[0].UserID,album_review[0].AlbumID).Find(&album).Error;err!=nil{
		log.Println("AddReviewsToAlbum2 "+err.Error())
		return err
	}

	if err:=Db.Self.Table("album").Where("user_id=? AND album_id=?",album.UserID,album.AlbumID).Update("content_sum",album.ContentSum-len(album_review)).Error; err!=nil{
		log.Println("AddReviewsToAlbum3 "+err.Error())
		return err
	}
	return nil
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


