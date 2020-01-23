//存放review的结构
package model

type Review struct {
    title string
    content string
    tag string
    picture string
    comment_sum string
    like_sum string
}

/*type Count struct {
    count string
}*/

type UserInfo struct {
    name string
    user_picture string
}

type CommentInfo struct {
    user_id string
    name string
    user_picture string
    comment_id string
    content string
    time string
    like_sum string
}

//获取likesum
type Sum struct {
    sum int
}
