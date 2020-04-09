//存放review的结构
package model

type Review struct {
	Title       string
	Content     string
	Tag         string
	Picture     string
	Comment_sum int
	Like_sum    int
}

/*type Count struct {
    count string
}*/

//删掉了user_picture字段
type UserInfo struct {
	Name string
}

type CommentInfo struct {
	User_id      string
	Name         string
	User_picture string
	Comment_id   int
	Content      string
	Time         string
	Like_sum     int
	Comment_like bool
}

//获取likesum
type Sum struct {
	Sum int
}

type Comment struct {
	Content  string
	Like_sum int
}

type CommentID struct {
	Comment_id int
}

type ReviewID struct {
	Review_id int
}

type ReviewLike struct {
	Review_id int
	User_id   string
}

type LikeSum struct {
	Like_sum int
}

type Collection struct {
	User_id   string
	Review_id int
}

type CommentSum struct {
	Comment_sum int
}

type CommentLike struct {
	User_id    string
	Comment_id int
	Review_id  int
}

type CommentLikeSum struct {
	Like_sum int
}

type UserID struct {
	User_id string
}
