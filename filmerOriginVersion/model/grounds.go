//存储发现页面的结构
package model

//要返回的页面
type GroundInfos struct {
	User_id      string
	Name         string
	User_picture string
	Review_id    string
	Title        string
	Content      string
	Time         string
	Tag          string
	Picture      string
	Comment_sum  string
	Like_sum     string
}

//用来导出关注的user
type GroundInfosID struct {
	User_id2     string
	Name         string
	User_picture string
	Review_id    string
	Title        string
	Content      string
	Time         string
	Tag          string
	Picture      string
	Comment_sum  string
	Like_sum     string
}

//用来计算轮到哪个评论区间
var Count int

//用来计算评论的总数
var CountSum int
