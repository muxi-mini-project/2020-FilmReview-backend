//存放review的结构
package model

type Review struct {
	title       string
	content     string
	tag         string
	picture     string
	comment_sum int
	like_sum    int
}

/*type Count struct {
    count string
}*/

type UserInfo struct {
	name         string
	user_picture string
}

type CommentInfo struct {
	user_id      string
	name         string
	user_picture string
	comment_id   string
	content      string
	time         string
	like_sum     int
}

//获取likesum
type Sum struct {
	sum int
}

type Comment struct {
	content  string
	like_sum int
}

type CommentID struct {
	comment_id int
}

type ReviewID struct {
	review_id int
}
