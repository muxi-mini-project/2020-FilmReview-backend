//存放创建review的函数
package Func

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/database"
	"github.com/muxi-mini-project/2020-FilmReview-backend/filmer/model"
	"log"
	"strconv"
	"sync"
)

//通过全局变量CountSum实现
/*func CountRecord() (string,error) {
    var count Count
    sql := "select count(*) from review"
    err:= database.DB.Raw(sql).Scan(&count).Error()

    count.count=strconv.Atoi(count.count)
    return stronv.Itoa(100000+count.Count),err
}*/

func InsertReview(review model.Review, reviewID string, userInfo mdoel.UserInfo, userid string) error {
	sql := "insert into review(user_id,name,user_picture,review_id,title,content,time,tag,picture,comment_sum,like_sum) values(" + userid + "," + userInfo.name + "," + userInfo.user_picture + "," + reviewID + "," + review.title + "," + review.content + "," + review.tag + "," + review.picture + "," + "," + comment_sum + "," + review.like_sum + ")"
	err := database.DB.Raw(sql).Error()
	return err
}

func GetUserInfo(userID string) (model.UserInfo, error) {
	var userInfo model.UserInfo
	sql := "select name,user_picture,from user where user_id = " + userID

	if err := database.DB.Raw(sql).Scan(&userInfo).Error(); err != nil {
		return errors.New("server busy")
	}
	return nil
}

func GetReview(reviewID string) ([]model.CommentInfo, error) {
	var comment []model.CommentInfo
	sql := "select user_id,name,user_picture,comment_id,content,time,like_sum from comment where review_id = " + reviewID
	if err := database.DB.Raw(sql).Scan(&comment).Error(); err != nil {
		return nil, errors.New("surver busy")
	}
	return comment, nil
}

//还有错误处理
func GetReviewCollection(reviewID, userID string) bool {
	sql := "select user_id from collection where user_id = " + userID + " AND " + "review_id = " + reviewID
	if ok := database.DB.Raw(sql).RecordNotFound; ok {
		return false
	}

	return true
}

func GetReviewLike(reviewID, userID string) bool {
	sql := "select user_id from review_like where user_id = " + userID + " AND " + "review_id = " + reviewID
	if ok := database.DB.Raw(sql).RecordNotFound; ok {
		return false
	}

	return true
}

func GetCommentLike(commentID, userID string) bool {
	sql := "select user_id from comment_like where user_id =" + userID + " AND " + "comment_id = " + commentID
	if ok := database.DB.Raw(sql).RecordNotFound; ok {
		return false
	}
	return true
}

func GetExtraInfo(comment []model.CommentInfo, userID string, reviewID string) (bool, bool, bool) {
	var com, col, rev bool
	wg := sync.Mutex
	wg.Add(len(comment) + 2)

	go func() {
		col = GetReviewCollection(reviewID, userID)
		wg.Done
	}()

	go func() {
		rev = GetReviewLike(reviewID, userID)
		wg.Done
	}()

	for i := 0; i < len(comment); i++ {
		go func() {
			com = GetCommentLike(comment[i].Comment_id, userID)
			wg.Done
		}()
	}

	wg.Wait()

	return com, rev, col
}

func DeleteReview(userID, reviewID string) error {
	sql := "delete from user_review where user_id = " + userID + " AND review_id = " + reviewID
	err := database.DB.Raw(sql).Error()
	return err
}

func DeleteReviewCollection(reviewID string) error {
	sql := "delete from collection where review_id = " + reviewID
	err := database.DB.Raw(sql).Error()
	return err
}

func DeleteReviewLike(reviewID string) error {
	sql := "delete from review_like where review_id = " + reviewID
	err := database.DB.Raw(sql).Error()
	return err
}

func DeleteReviewComment(reviewID string) error {
	sql := "delete from comment where review_id = " + reviewID
	err := database.DB.Raw(sql).Error()
	return err
}

func DeleteReviewCommentLike(reviewID string) error {
	sql := "delete from comment_like where review_id = " + reviewID
	err := database.DB.Raw(sql).Error()
	return err
}

func DeleteAlbumReview(reviewID string) error {
	sql := "delete from album_review where review_id = " + reviewID
	err := database.DB.Raw(sql).Error()
	return err
}

func DeleteFunc(userID, reviewID string) error {
	wg := sync.WaitGroup{}
	wg.Add(6)

	go func() {
		if err := DeleteReview(userID, reviewID); err != nil {
			return errors.New("surver bussy")
		}
		wg.Done
	}()

	go func() {
		if err := DeleteReviewCollection(reviewID); err != nil {
			return errors.New("surver bussy")
		}
		wg.Done
	}()

	go func() {
		if err := DeleteReviewLike(reviewID); err != nil {
			return errors.New("surver bussy")
		}
		wg.Done
	}()

	go func() {
		if err := DeleteReviewComment(reviewID); err != nil {
			return errors.New("surver bussy")
		}
		wg.Done
	}()

	go func() {
		if err := DeleteReviewCommentLike(reviewID); err != nil {
			return errors.New("surver bussy")
		}
		wg.Done
	}()

	go func() {
		if err := DeleteAlbumReview(reviewID); err != nil {
			return errors.New("surver bussy")
		}
		wg.Done
	}()

	wg.Wait()
	return nil
}

//错误
func ChangeReviewLikeFunc(userID, reviewID string) {
	sql := "select *from review_like where user_id = " + userID + " AND review_id = " + reviewID
	if ok := database.DB.Raw(sql).RecordNotFound(); ok {
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			sql1 := "insert into review_like(user_id,review_id) values(" + userID + "," + reviewID + ")"
			database.DB.Raw(sql1)
			wg.Done()
		}()

		go func() {
			AddReviewLikeSum(reviewID)
			wg.Done()
		}()

		wg.Wait()

		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		sql1 := "delete from review_like where user_id = " + userID + " AND review_id = " + reviewID
		database.DB.Raw(sql1)
		wg.Done()
	}()

	go func() {
		DeleteReviewLikeSum(reviewID)
		wg.Done()
	}()

	wg.Wait()
	return
}

//修改like_sum
func AddReviewLikeSum(reviewID string) {
	var likeSum model.Sum
	sql := "select like_sum from review where review_id = " + reviewID
	database.DB.Raw(sql).Scan(&likeSum)
	likeSum += 1
	sql2 := "update review set like_sum = " + strconv.Atoi(likeSum) + " where review_id = " + reviewID
	database.DB.Raw(sql2)
	return
}

func DeleteReviewLikeSum(reviewID string) {
	var likeSum model.Sum
	sql := "select like_sum from review where review_id = " + reviewID
	database.DB.Raw(sql).Scan(&likeSum)
	likeSum -= 1
	sql2 := "update review set like_sum = " + strconv.Atoi(likeSum) + " where review_id = " + reviewID
	database.DB.Raw(sql2)
	return
}

func NewCollection(userID, reviewID string) {
	sql := "select *from collection where user_id = " + userID + " AND reviewID = " + reviewID
	if ok := database.DB.Raw(sql).RecordNotFound(); ok {
		sql1 := "insert into collection(user_id,review_id) values(" + userID + "," + reviewID + ")"
		database.DB.Raw(sql1)
		return
	}

	sql1 := "delete from collection where user_id = " + userID + " AND review_id = " + reviewID
	database.DB.Raw(sql1)
	return
}

func GetCommentID() int {
	var commentID model.CommentID
	sql := "select count(*) from comment"
	database.Raw(sql).Scan(&commentID)
	return commentID
}

func InsertComment(userID string, userInfo model.UserInfo, reviewID string, comment model.Comment, commentID int) error {
	sql := "insert into comment(user_id,name,user_picture,review_id,comment_id,content,time,like_sum) values(" + userID + "," + userInfo.name + "," + userInfo.user_picture + "," + reviewID + "," + commentID + "," + comment.content + "," + comment.like_sum + ")"
	if err := database.Raw(sql).Error(); err != nil {
		return errors.New("surver busy")
	}

	return nil
}

//修改评论数
func AddCommentSum(reviewID string) {
	var commentSum model.Sum
	sql := "select comment_sum from review where review_id = " + reviewID
	database.DB.Raw(sql).Scan(&commentSum)
	likeSum += 1
	sql2 := "update review set comment_sum = " + strconv.Atoi(likeSum) + " where review_id = " + reviewID
	database.DB.Raw(sql2)
	return
}

func DeleteCommentSum(reviewID string) {
	var commentSum model.Sum
	sql := "select comment_sum from review where review_id = " + reviewID
	database.DB.Raw(sql).Scan(&commentSum)
	likeSum -= 1
	sql2 := "update review set comment_sum = " + strconv.Atoi(likeSum) + " where review_id = " + reviewID
	database.DB.Raw(sql2)
	return
}

func NewComment(userID string, reviewID string, comment model.Comment) error {
	//并发处理获取commentid和user信息
	wg := sync.WaitGroup{}
	wg.Add(2)

	var commentID int
	var userInfo model.UserInfo
	var err1 error

	go func() {
		//获取commentid
		commentID = GetCommentID()
		commentID++
		wg.Done()
	}()

	go func() {
		//获取userid的信息
		userInfo, err1 = GetUserInfo(userID)
		if err1 != nil {
			return errors.New("server busy")
		}
		wg.Done()
	}()

	wg.Wait()

	//再并发处理插入comment和修改commentsum

	wg := sync.WaitGroup{}
	wg.Add(2)

	var err2 error

	go func() {
		//插入comment表
		if err2 = Func.InsertComment(userID, userInfo, reviewID, comment, commentID); err2 != nil {
			return errors.New("server busy")
		}
		wg.Done()
	}()

	go func() {
		//修改review的评论总数
		AddCommentSum := Func.CommentSum(reviewID)
		wg.Done()
	}()

	wg.Wait()

	return nil
}

func NewCommentLike(userID, commentID string) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var ok bool
	var reviewID model.ReviewID

	go func() {
		sql := "select *from comment_like where user_id = " + userID + " AND comment_id = " + commentID
		ok := database.DB.Raw(sql).RecordNotFound()
		wg.Done()

	}()

	go func() {
		sql := "select review_id from comment where comment_id = " + commentID
		database.Raw(sql).Scan(reviewID)
		wg.Done()
	}()

	wg.Wait()

	if ok {
		sql1 := "insert into comment_like(user_id,comment_id,review_id) values(" + userID + "," + commentID + "," + reviewID.review_id + ")"
		database.DB.Raw(sql1)
		return
	}

	sql1 := "delete from collection where user_id = " + userID + " AND comment_id = " + commentID
	database.DB.Raw(sql1)
	return
}

func DeleteComment(commentID int) {
	sql := "delete from comment where comment_id = " + commentID
	database.DB.Raw(sql)
	return
}

func DeleteCommentLike(commentID int) {
	sql := "delete from comment_like where comment_id = " + commentID
	database.DB.Raw(sql)
	return
}

func DeleteCommentFunc(commentID int) {
	var reviewID model.ReviewID
	sql := "select review_id from comment where comment_id = " + commentID
	database.Raw(sql).Scan(reviewID)
	wg := sync.Wait{}
	wg.Add(3)

	go func() {
		DeleteComment(commentID)
		wg.Done()
	}()

	go func() {
		DeleteCommentLike(commentID)
		wg.Done()
	}()

	go func() {
		DeleteCommentSum(reviewID.review_id)
		wg.Done()
	}()

	wg.Wait()

	return
}
