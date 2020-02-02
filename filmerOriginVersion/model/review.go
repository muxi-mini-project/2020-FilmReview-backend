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

type UserInfo struct {
	Name         string
	User_picture string
}

type CommentInfo struct {
	User_id      string
	Name         string
	User_picture string
	Comment_id   string
	Content      string
	Time         string
	Like_sum     int
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
	Review_id string
}
