//存放创建review的函数
package Func

import (
	"fmt"
	"log"
	//"github.com/gin-gonic/gin"
	"github.com/filmer2/database"
	"github.com/filmer2/model"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	//"strconv"
	"errors"
	//"sync"
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
	sql := "insert into user_review(user_id,name,review_id,title,content,tag,picture,comment_sum,like_sum) values"
	sql += fmt.Sprintf("('%s','%s',%d,'%s','%s','%s','%s',%d,%d);", userid, userInfo.Name, reviewID, review.Title, review.Content, review.Tag, review.Picture, 0, 0)
	log.Println(sql)
	err := database.DB.Exec(sql).Error
	log.Println("InsertReview")
	log.Println(err)
	return err
}

func GetUserInfo(userID string) (model.UserInfo, error) {
	log.Println(userID)
	var userInfo model.UserInfo
	sql := "select name from user where user_id = '" + userID + "';"

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
func GetReviewCollection(reviewID int, userID string) (error, bool) {
	var userid model.UserID
	sql := "select user_id from collection where user_id = " + userID + " AND " + "review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	if err := database.DB.Raw(sql).Scan(&userid).Error; err != nil {
		return err, false
	}
	if userid.User_id == "" {
		return nil, false
	}

	return nil, true
}

func GetReviewLike(reviewID int, userID string) (error, bool) {
	var userid model.UserID
	sql := "select user_id from review_like where user_id = " + userID + " AND " + "review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	if err := database.DB.Raw(sql).Scan(&userid).Error; err != nil {
		return err, false
	}
	if userid.User_id == "" {
		return nil, false
	}

	return nil, true
}

func GetCommentLike(userID string, comment *[]model.CommentInfo) error {
	for i := 0; i < len((*comment)); i++ {
		log.Println(i, "start")
		var userid model.UserID
		log.Println("userid done")
		sql := "select user_id from comment_like where user_id =" + userID + " AND " + "comment_id = "
		sql += fmt.Sprintf("%d", (*comment)[i].Comment_id)
		if err := database.DB.Raw(sql).Scan(&userid).Error; err != nil {
			return err
		}
		if userid.User_id == "" {
			log.Println("false func")
			(*comment)[i].Comment_like = false
			return nil
		}
		log.Println("true func")
		(*comment)[i].Comment_like = true
		fmt.Println(i, "Done")
	}
	return nil
}

func GetExtraInfo(comment *[]model.CommentInfo, userID string, reviewID int) (error,bool, bool) {
	var col, rev bool
	var err1, err2, err3 error
	log.Println(len((*comment)))
	//	log.Println((*comment)[1], (*comment)[0])
	errChannel := make(chan error, 3)
	defer close(errChannel)

	go func() {
		err1, col = GetReviewCollection(reviewID, userID)
		errChannel <- err1
	}()

	go func() {
		err2,rev = GetReviewLike(reviewID, userID)
		errChannel <- err2
	}()
	go func() {
		err3=GetCommentLike(userID, comment)
		errChannel <- err3
	}()

	for i := 0; i < 3; i++ {
		if err:=<-errChannel;err!=nil {
			return errors.New("Get Extra Fail"), false, false
		}
	}

	return nil, rev, col
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

	errChannel := make(chan error, 6)
	defer close(errChannel)

	go func() {
		err := DeleteReview(userID, reviewID)
		errChannel <- err
	}()

	go func() {
		err := DeleteReviewCollection(reviewID)
		errChannel <- err
	}()

	go func() {
		err := DeleteReviewLike(reviewID)
		errChannel <- err
	}()

	go func() {
		err := DeleteReviewComment(reviewID)
		errChannel <- err
	}()

	go func() {
		err := DeleteReviewCommentLike(reviewID)
		errChannel <- err
	}()

	go func() {
		err := DeleteAlbumReview(reviewID)
		errChannel <- err
	}()

	for i := 0; i < 6; i++ {
		if err := <-errChannel; err != nil {
			return err
		}
	}

	return nil
}

//错误
func ChangeReviewLikeFunc(userID string, reviewID int) error {
	sql := "select *from review_like where user_id = " + userID + " AND review_id = "
	sql += fmt.Sprintf("%d", reviewID)
	var reviewlike model.ReviewLike
	if err := database.DB.Raw(sql).Scan(&reviewlike).Error; err != nil {
		return err
	}
	log.Println(reviewlike)
	if reviewlike.User_id == "" {
		log.Println("AddFunc")
		errChannel := make(chan error, 2)
		defer close(errChannel)

		go func() {
			sql1 := "insert into review_like(user_id,review_id) values(" + userID + ","
			sql1 += fmt.Sprintf("%d);", reviewID)
			err := database.DB.Exec(sql1).Error
			errChannel <- err
		}()

		go func() {
			err := AddReviewLikeSum(reviewID)
			errChannel <- err
		}()

		for i := 0; i < 2; i++ {
			if err := <-errChannel; err != nil {
				return err
			}
		}

		return nil
	}

	log.Println("DeleteFunc")
	errChannel2 := make(chan error, 2)
	defer close(errChannel2)

	go func() {
		sql1 := "delete from review_like where user_id = " + userID + " AND review_id = "
		sql1 += fmt.Sprintf("%d;", reviewID)
		err := database.DB.Exec(sql1).Error
		errChannel2 <- err
	}()

	go func() {
		err := DeleteReviewLikeSum(reviewID)
		errChannel2 <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChannel2; err != nil {
			return err
		}
	}
	return nil
}

//修改like_sum
func AddReviewLikeSum(reviewID int) error {
	var likeSum model.LikeSum
	sql := "select like_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d;", reviewID)
	if err := database.DB.Raw(sql).Scan(&likeSum).Error; err != nil {
		return err
	}
	log.Println(likeSum)
	likeSum.Like_sum += 1
	sql2 := "update user_review set like_sum = "
	sql2 += fmt.Sprintf("%d where review_id = %d ;", likeSum.Like_sum, reviewID)
	if err2 := database.DB.Exec(sql2).Error; err2 != nil {
		return err2
	}
	return nil
}

func DeleteReviewLikeSum(reviewID int) error {
	var likeSum model.LikeSum
	sql := "select like_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	if err := database.DB.Raw(sql).Scan(&likeSum).Error; err != nil {
		return err
	}
	log.Println(likeSum)
	likeSum.Like_sum -= 1
	sql2 := "update user_review set like_sum = "
	sql2 += fmt.Sprintf("%d where review_id = %d ;", likeSum.Like_sum, reviewID)
	if err2 := database.DB.Exec(sql2).Error; err2 != nil {
		return err2
	}
	return nil
}

func NewCollection(userID string, reviewID int) error {
	sql := "select *from collection where user_id = " + userID + " AND review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	var collection model.Collection
	if err := database.DB.Raw(sql).Scan(&collection).Error; err != nil {
		return err
	}
	if collection.User_id == "" {
		sql1 := "insert into collection(user_id,review_id) values(" + userID + ","
		sql1 += fmt.Sprintf("%d );", reviewID)
		if err2 := database.DB.Exec(sql1).Error; err2 != nil {
			return err2
		}
		return nil
	}

	sql1 := "delete from collection where user_id = " + userID + " AND review_id = "
	sql1 += fmt.Sprintf("%d ;", reviewID)
	if err3 := database.DB.Exec(sql1).Error; err3 != nil {
		return err3
	}
	return nil
}

func GetCommentID() (error, int) {
	var commentID model.CommentID
	sql := "select max(comment_id) as comment_id from comment"
	if err := database.DB.Raw(sql).Scan(&commentID).Error; err != nil {
		return err, 0
	}
	return nil, commentID.Comment_id
}

func InsertComment(userID string, userInfo model.UserInfo, reviewID int, comment model.Comment, commentID int) error {
	sql := "insert into comment(user_id,name,review_id,comment_id,content,like_sum) "
	sql += fmt.Sprintf(" values('%s','%s',%d,%d,'%s',%d);", userID, userInfo.Name, reviewID, commentID, comment.Content, 0)
	if err := database.DB.Exec(sql).Error; err != nil {
		return errors.New("server busy")
	}

	return nil
}

//修改评论数
func AddCommentSum(reviewID int) error {
	var commentSum model.CommentSum
	sql := "select comment_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	if err := database.DB.Raw(sql).Scan(&commentSum).Error; err != nil {
		return err
	}
	commentSum.Comment_sum += 1
	sql2 := "update user_review set comment_sum = "
	sql2 += fmt.Sprintf(" %d where review_id = %d ;", commentSum.Comment_sum, reviewID)
	if err2 := database.DB.Exec(sql2).Error; err2 != nil {
		return err2
	}
	return nil
}

func DeleteCommentSum(reviewID int) error {
	var commentSum model.CommentSum
	sql := "select comment_sum from user_review where review_id = "
	sql += fmt.Sprintf("%d ;", reviewID)
	if err := database.DB.Raw(sql).Scan(&commentSum).Error; err != nil {
		return err
	}
	commentSum.Comment_sum -= 1
	sql2 := "update user_review set comment_sum = "
	sql2 += fmt.Sprintf("%d where review_id = %d", commentSum.Comment_sum, reviewID)
	if err2 := database.DB.Exec(sql2).Error; err2 != nil {
		return err2
	}
	return nil
}

func NewComment(userID string, reviewID int, comment model.Comment) (error, int) {
	//并发处理获取commentid和user信息
	errChannel := make(chan error, 2)
	defer close(errChannel)

	var commentID int
	var userInfo model.UserInfo
	var err, err1 error

	go func() {
		//获取commentid
		err, commentID = GetCommentID()
		commentID++
		errChannel <- err
	}()

	go func() {
		//获取userid的信息
		userInfo, err1 = GetUserInfo(userID)
		errChannel <- err1
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChannel; err != nil {
			return err, 0
		}
	}

	//再并发处理插入comment和修改commentsum

	errChannel2 := make(chan error, 2)
	defer close(errChannel2)

	var err2 error

	go func() {
		//插入comment表
		err2 = InsertComment(userID, userInfo, reviewID, comment, commentID)
		errChannel2 <- err2
	}()

	go func() {
		//修改review的评论总数
		err := AddCommentSum(reviewID)
		errChannel2 <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChannel; err != nil {
			return err, 0
		}
	}

	return nil, commentID
}

func AddComment_Like(commentID int) error {
	var commentlikesum model.CommentLikeSum
	sql := "select like_sum from comment where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	if err := database.DB.Raw(sql).Scan(&commentlikesum).Error; err != nil {
		return err
	}
	commentlikesum.Like_sum += 1
	sql2 := "update comment set like_sum = "
	sql2 += fmt.Sprintf(" %d where comment_id = %d ;", commentlikesum.Like_sum, commentID)
	if err2 := database.DB.Exec(sql2).Error; err2 != nil {
		return err2
	}
	return nil
}

func DeleteComment_Like(commentID int) error {
	var commentlikesum model.CommentLikeSum
	sql := "select like_sum from comment where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	if err := database.DB.Raw(sql).Scan(&commentlikesum).Error; err != nil {
		return err
	}
	commentlikesum.Like_sum -= 1
	sql2 := "update comment set like_sum = "
	sql2 += fmt.Sprintf(" %d where comment_id = %d ;", commentlikesum.Like_sum, commentID)
	if err2 := database.DB.Exec(sql2).Error; err2 != nil {
		return err2
	}
	return nil
}

func NewCommentLike(userID string, commentID int) error {
	errChannel := make(chan error, 2)

	defer close(errChannel)
	var commentLike model.CommentLike
	var reviewID model.ReviewID

	go func() {
		sql := "select *from comment_like where user_id = " + userID + " AND comment_id = "
		sql += fmt.Sprintf("%d", commentID)
		err := database.DB.Raw(sql).Scan(&commentLike).Error
		errChannel <- err

	}()

	go func() {
		sql := "select review_id from comment where comment_id = "
		sql += fmt.Sprintf("%d ;", commentID)
		err := database.DB.Raw(sql).Scan(&reviewID).Error
		errChannel <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChannel; err != nil {
			return err
		}
	}

	if commentLike.User_id == "" {
		errChannel2 := make(chan error, 2)
		defer close(errChannel2)

		go func() {
			sql1 := "insert into comment_like(user_id,comment_id,review_id) "
			sql1 += fmt.Sprintf("values('%s',%d,%d);", userID, commentID, reviewID.Review_id)
			err := database.DB.Exec(sql1).Error
			errChannel2 <- err
		}()

		go func() {
			err := AddComment_Like(commentID)
			errChannel2 <- err
		}()

		for i := 0; i < 2; i++ {
			if err := <-errChannel2; err != nil {
				return err
			}
		}
		return nil
	}

	errChannel3 := make(chan error, 2)
	defer close(errChannel3)

	go func() {
		sql1 := "delete from comment_like where user_id = " + userID + " AND comment_id = "
		sql1 += fmt.Sprintf("%d ;", commentID)
		err := database.DB.Exec(sql1).Error
		errChannel3 <- err
	}()

	go func() {
		err := DeleteComment_Like(commentID)
		errChannel3 <- err
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChannel3; err != nil {
			return err
		}
	}
	return nil
}

func DeleteComment(commentID int) error {
	sql := "delete from comment where "
	sql += fmt.Sprintf("comment_id = %d ;", commentID)
	if err := database.DB.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCommentLike(commentID int) error {
	sql := "delete from comment_like where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	if err := database.DB.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCommentFunc(commentID int) error {
	var reviewID model.ReviewID
	sql := "select review_id from comment where comment_id = "
	sql += fmt.Sprintf("%d ;", commentID)
	database.DB.Raw(sql).Scan(&reviewID)
	errChannel := make(chan error, 3)
	defer close(errChannel)

	go func() {
		err := DeleteComment(commentID)
		errChannel <- err
	}()

	go func() {
		err := DeleteCommentLike(commentID)
		errChannel <- err
	}()

	go func() {
		err := DeleteCommentSum(reviewID.Review_id)
		errChannel <- err
	}()

	for i := 0; i < 3; i++ {
		if err := <-errChannel; err != nil {
			return err
		}
	}

	return nil
}
