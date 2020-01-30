//存储发现页面的结构
package model

//要返回的页面
type GroundInfos struct {
	user_id      string
	name         string
	user_picture string
	review_id    string
	title        string
	content      string
	time         string
	tag          string
	picture      string
	comment_sum  string
	like_sum     string
}

//用来导出关注的user
type GroundInfosID struct {
	user_id2     string
	name         string
	user_picture string
	review_id    string
	title        string
	content      string
	time         string
	tag          string
	picture      string
	comment_sum  string
	like_sum     string
}

//用来计算轮到哪个评论区间
var Count int

//用来计算评论的总数
var CountSum int
