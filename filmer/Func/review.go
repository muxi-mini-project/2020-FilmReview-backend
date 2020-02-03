//存放创建review的函数
package Func

import (
	"errors"
	"fmt"
	"log"
	//"github.com/gin-gonic/gin"
	"github.com/filmer/database"
	"github.com/filmer/model"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	//"strconv"
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

func InsertReview(review model.Review, reviewID int, userInfo model.UserInfo, userid string) error {
	sql := "insert into user_review(user_id,name,user_picture,review_id,title,content,tag,picture,comment_sum,like_sum) values"
	sql += fmt.Sprintf("('%s','%s','%s',%d,'%s','%s','%s','%s',%d,%d);", userid, userInfo.Name, userInfo.User_picture, reviewID, review.Title, review.Content, review.Tag, review.Picture, 0, 0)
	log.Println(sql)
	err := database.DB.Exec(sql).Error
	log.Println("InsertReview")
	log.Println(err)
	return err
}

func GetUserInfo(userID string) (model.UserInfo, error) {
	var userInfo model.UserInfo
	sql := "select name,user_picture from user where user_id = " + userID

	//Raw用来查询，Exec用来修改插入
	if err := database.DB.Raw(sql).Scan(&userInfo).Error; err != nil {
		return userInfo, errors.New("server busy")
	}
	log.Println("GetUserInfo")
	log.Println(userInfo)
	return userInfo, nil
}

func GetReview(reviewID int) ([]model.CommentInfo, error) {
	var comment []model.CommentInfo
	sql := "select user_id,name,user_picture,comment_id,content,time,like_sum from comment where review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	if err := database.DB.Raw(sql).Scan(&comment).Error; err != nil {
		return nil, errors.New("surver busy")
	}
	log.Println(comment)
	return comment, nil
}

//还有错误处理
func GetReviewCollection(reviewID int, userID string) bool {
	var userid model.UserID
	sql := "select user_id from collection where user_id = " + userID + " AND " + "review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	database.DB.Raw(sql).Scan(&userid)
	if userid.User_id == "" {
		return false
	}

	return true
}

func GetReviewLike(reviewID int, userID string) bool {
	var userid model.UserID
	sql := "select user_id from review_like where user_id = " + userID + " AND " + "review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	database.DB.Raw(sql).Scan(&userid)
	if userid.User_id == "" {
		return false
	}

	return true
}

func GetCommentLike(userID string, comment *[]model.CommentInfo) {
	for i := 0; i < len((*comment)); i++ {
		log.Println(i, "start")
		var userid model.UserID
		log.Println("userid done")
		sql := "select user_id from comment_like where user_id =" + userID + " AND " + "comment_id = "
		sql += fmt.Sprintf("%d", (*comment)[i].Comment_id)
		database.DB.Raw(sql).Scan(&userid)
		if userid.User_id == "" {
			log.Println("false func")
			(*comment)[i].Comment_like = false
			return
		}
		log.Println("true func")
		(*comment)[i].Comment_like = true
		fmt.Println(i, "Done")
	}
	return
}

func GetExtraInfo(comment *[]model.CommentInfo, userID string, reviewID int) (bool, bool) {
	var col, rev bool
	log.Println(len((*comment)))
	log.Println((*comment)[1], (*comment)[0])
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		col = GetReviewCollection(reviewID, userID)
		wg.Done()
	}()

	go func() {
		rev = GetReviewLike(reviewID, userID)
		wg.Done()
	}()
	go func() {
		GetCommentLike(userID, comment)
		wg.Done()
	}()

	wg.Wait()

	return rev, col
}

func DeleteReview(userID string, reviewID int) error {
	sql := "delete from user_review where user_id = " + userID + " AND review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	err := database.DB.Exec(sql).Error
	return err
}

func DeleteReviewCollection(reviewID int) error {
	sql := "delete from collection where review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	err := database.DB.Exec(sql).Error
	return err
}

func DeleteReviewLike(reviewID int) error {
	sql := "delete from review_like where review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	err := database.DB.Exec(sql).Error
	return err
}

func DeleteReviewComment(reviewID int) error {
	sql := "delete from comment where review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	err := database.DB.Exec(sql).Error
	return err
}

func DeleteReviewCommentLike(reviewID int) error {
	sql := "delete from comment_like where review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	err := database.DB.Exec(sql).Error
	return err
}

func DeleteAlbumReview(reviewID int) error {
	sql := "delete from album_review where review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	err := database.DB.Exec(sql).Error
	return err
}

func DeleteFunc(userID string, reviewID int) error {
	wg := sync.WaitGroup{}
	wg.Add(6)

	go func() {
		if err := DeleteReview(userID, reviewID); err != nil {
			//return errors.New("surver bussy")
		}
		wg.Done()
	}()

	go func() {
		if err := DeleteReviewCollection(reviewID); err != nil {
			//return errors.New("surver bussy")
		}
		wg.Done()
	}()

	go func() {
		if err := DeleteReviewLike(reviewID); err != nil {
			//return errors.New("surver bussy")
		}
		wg.Done()
	}()

	go func() {
		if err := DeleteReviewComment(reviewID); err != nil {
			//return errors.New("surver bussy")
		}
		wg.Done()
	}()

	go func() {
		if err := DeleteReviewCommentLike(reviewID); err != nil {
			//return errors.New("surver bussy")
		}
		wg.Done()
	}()

	go func() {
		if err := DeleteAlbumReview(reviewID); err != nil {
			//return errors.New("surver bussy")
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

//错误
func ChangeReviewLikeFunc(userID string, reviewID int) {
	sql := "select *from review_like where user_id = " + userID + " AND review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	var reviewlike model.ReviewLike
	database.DB.Raw(sql).Scan(&reviewlike)
	log.Println(reviewlike)
	if reviewlike.User_id == "" {
		log.Println("AddFunc")
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			sql1 := "insert into review_like(user_id,review_id) values(" + userID + ","
			sql1 += fmt.Sprintf("%d);", reviewID)
			database.DB.Exec(sql1)
			wg.Done()
		}()

		go func() {
			AddReviewLikeSum(reviewID)
			wg.Done()
		}()

		wg.Wait()

		return
	}

	log.Println("DeleteFunc")
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		sql1 := "delete from review_like where user_id = " + userID + " AND review_id = "
		sql1 += fmt.Sprintf("%d;", reviewID)
		database.DB.Exec(sql1)
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
func AddReviewLikeSum(reviewID int) {
	var likeSum model.LikeSum
	sql := "select like_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d;", reviewID)
	database.DB.Raw(sql).Scan(&likeSum)
	log.Println(likeSum)
	likeSum.Like_sum += 1
	sql2 := "update user_review set like_sum = "
	sql2 += fmt.Sprintf("%d where review_id = %d ;", likeSum.Like_sum, reviewID)
	database.DB.Exec(sql2)
	return
}

func DeleteReviewLikeSum(reviewID int) {
	var likeSum model.LikeSum
	sql := "select like_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	database.DB.Raw(sql).Scan(&likeSum)
	log.Println(likeSum)
	likeSum.Like_sum -= 1
	sql2 := "update user_review set like_sum = "
	sql2 += fmt.Sprintf("%d where review_id = %d ;", likeSum.Like_sum, reviewID)
	database.DB.Exec(sql2)
	return
}

func NewCollection(userID string, reviewID int) {
	sql := "select *from collection where user_id = " + userID + " AND review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	var collection model.Collection
	database.DB.Raw(sql).Scan(&collection)
	if collection.User_id == "" {
		sql1 := "insert into collection(user_id,review_id) values(" + userID + ","
		sql1 += fmt.Sprintf("%d );", reviewID)
		database.DB.Exec(sql1)
		return
	}

	sql1 := "delete from collection where user_id = " + userID + " AND review_id = "
	sql1 += fmt.Sprintf("%d ;", reviewID)
	database.DB.Exec(sql1)
	return
}

func GetCommentID() int {
	var commentID model.CommentID
	sql := "select max(comment_id) as comment_id from comment"
	database.DB.Raw(sql).Scan(&commentID)
	return commentID.Comment_id
}

func InsertComment(userID string, userInfo model.UserInfo, reviewID int, comment model.Comment, commentID int) error {
	sql := "insert into comment(user_id,name,user_picture,review_id,comment_id,content,like_sum) "
	sql += fmt.Sprintf(" values('%s','%s','%s',%d,%d,'%s',%d);", userID, userInfo.Name, userInfo.User_picture, reviewID, commentID, comment.Content, 0)
	if err := database.DB.Exec(sql).Error; err != nil {
		return errors.New("surver busy")
	}

	return nil
}

//修改评论数
func AddCommentSum(reviewID int) {
	var commentSum model.CommentSum
	sql := "select comment_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	database.DB.Raw(sql).Scan(&commentSum)
	commentSum.Comment_sum += 1
	sql2 := "update user_review set comment_sum = "
	sql2 += fmt.Sprintf(" %d where review_id = %d ;", commentSum.Comment_sum, reviewID)
	database.DB.Exec(sql2)
	return
}

func DeleteCommentSum(reviewID int) {
	var commentSum model.CommentSum
	sql := "select comment_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	database.DB.Raw(sql).Scan(&commentSum)
	commentSum.Comment_sum -= 1
	sql2 := "update user_review set comment_sum = "
	sql2 += fmt.Sprintf("%d where review_id = %d", commentSum.Comment_sum, reviewID)
	database.DB.Exec(sql2)
	return
}

func NewComment(userID string, reviewID int, comment model.Comment) (error, int) {
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
			//return errors.New("server busy")
		}
		wg.Done()
	}()

	wg.Wait()

	//再并发处理插入comment和修改commentsum

	wg2 := sync.WaitGroup{}
	wg2.Add(2)

	var err2 error

	go func() {
		//插入comment表
		if err2 = InsertComment(userID, userInfo, reviewID, comment, commentID); err2 != nil {
			//return errors.New("server busy")
		}
		wg2.Done()
	}()

	go func() {
		//修改review的评论总数
		AddCommentSum(reviewID)
		wg2.Done()
	}()

	wg2.Wait()

	return nil, commentID
}

func AddComment_Like(commentID int) {
	var commentlikesum model.CommentLikeSum
	sql := "select like_sum from comment where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	database.DB.Raw(sql).Scan(&commentlikesum)
	commentlikesum.Like_sum += 1
	sql2 := "update comment set like_sum = "
	sql2 += fmt.Sprintf(" %d where comment_id = %d ;", commentlikesum.Like_sum, commentID)
	database.DB.Exec(sql2)
}

func DeleteComment_Like(commentID int) {
	var commentlikesum model.CommentLikeSum
	sql := "select like_sum from comment where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	database.DB.Raw(sql).Scan(&commentlikesum)
	commentlikesum.Like_sum -= 1
	sql2 := "update comment set like_sum = "
	sql2 += fmt.Sprintf(" %d where comment_id = %d ;", commentlikesum.Like_sum, commentID)
	database.DB.Exec(sql2)
}

func NewCommentLike(userID string, commentID int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var commentLike model.CommentLike
	var reviewID model.ReviewID

	go func() {
		sql := "select *from comment_like where user_id = " + userID + " AND comment_id = "
		sql += fmt.Sprintf("%d", commentID)
		database.DB.Raw(sql).Scan(&commentLike)
		wg.Done()

	}()

	go func() {
		sql := "select review_id from comment where comment_id = "
		sql += fmt.Sprintf("%d ;", commentID)
		database.DB.Raw(sql).Scan(&reviewID)
		wg.Done()
	}()

	wg.Wait()

	if commentLike.User_id == "" {
		wg2 := sync.WaitGroup{}
		wg2.Add(2)

		go func() {
			sql1 := "insert into comment_like(user_id,comment_id,review_id) "
			sql1 += fmt.Sprintf("values('%s',%d,%d);", userID, commentID, reviewID.Review_id)
			database.DB.Exec(sql1)
			wg2.Done()
		}()

		go func() {
			AddComment_Like(commentID)
			wg2.Done()
		}()

		wg2.Wait()
		return
	}

	wg2 := sync.WaitGroup{}
	wg2.Add(2)

	go func() {
		sql1 := "delete from comment_like where user_id = " + userID + " AND comment_id = "
		sql1 += fmt.Sprintf("%d ;", commentID)
		database.DB.Exec(sql1)
		wg2.Done()
	}()

	go func() {
		DeleteComment_Like(commentID)
		wg2.Done()
	}()
	wg2.Wait()
	return
}

func DeleteComment(commentID int) {
	sql := "delete from comment where "
	sql += fmt.Sprintf("comment_id = %d ;", commentID)
	database.DB.Exec(sql)
	return
}

func DeleteCommentLike(commentID int) {
	sql := "delete from comment_like where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	database.DB.Exec(sql)
	return
}

func DeleteCommentFunc(commentID int) {
	var reviewID model.ReviewID
	sql := "select review_id from comment where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	database.DB.Raw(sql).Scan(&reviewID)
	wg := sync.WaitGroup{}
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
		DeleteCommentSum(reviewID.Review_id)
		wg.Done()
	}()

	wg.Wait()

	return
}
