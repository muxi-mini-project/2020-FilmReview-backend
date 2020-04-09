//存储发现页面的结构
package model

//要返回的页面
type GroundInfos struct {
	User_id      string
	Name         string
	User_picture string
	Review_id    int
	Title        string
	Content      string
	Time         string
	Tag          string
	Picture      string
	Comment_sum  int
	Like_sum     int
}

//用来导出关注的user
type GroundInfosID struct {
	User_id2     string
	Name         string
	User_picture string
	Review_id    int
	Title        string
	Content      string
	Time         string
	Tag          string
	Picture      string
	Comment_sum  int
	Like_sum     int
}

//用来计算轮到哪个评论区间
var Count int

//用来计算评论的总数
type CountSum struct {
	Countsum int
}

var Countsum CountSum

//记录上一次插入的reviewid
type LastReviewID struct {
	Review_id int
}

var Lastreviewid LastReviewID

type Follows struct {
	User_id2 string
}
