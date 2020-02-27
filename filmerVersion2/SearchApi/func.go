package SearchApi

import (
    "github.com/huichen/sego"
    "github.com/filmer/database"
    "github.com/filmer/model"
    _ "github.com/go-sql-driver/mysql"
    "sync"
    //"log"
)

func cutWords(sentence []byte) []string {
    var segmenter sego.Segmenter
    segmenter.LoadDictionary("/root/go/src/github.com/huichen/sego/data/dictionary.txt")
    segments := segmenter.Segment(sentence)
    words := sego.SegmentsToSlice(segments,false)
    return words
}

//先预处理查询没有重复的id，再返回
func doSearch(words []string,result *[]model.GroundInfosID) error {
    var sum int
    var m1 map[string]bool
    var lock sync.Mutex
    l := len(words)
    //上面两个通道固定是接收l个信号
    channelCount := make(chan int,l)
    errChannel := make(chan error,l)
    channelIndex := make(chan string,l*l)
    defer close(channelCount)
    defer close(errChannel)
    defer close(channelIndex)

    for i := 0;i < l;i++{
        go func() {
            var tags []tag
            sql := "select review_id from user_review where tag like '%"+words[i]+"%'"
            err := database.DB.Raw(sql).Scan(&tags).Error
            l := len(tags)
            errChannel <- err
            channelCount <- l
            for i := 0;i < l;i++{
                channelIndex <- tags[i].Tag
            }
        }()
    }

    for i := 0;i < l;i++{
        if <-errChannel != nil {
            return <-errChannel
        }
        sum += <-channelCount
    }

    for i := 0;i < sum ;i++{
        m1[<-channelIndex] = true
    }

    //去重结束，接下来可以一次性查找
    l2 := len(m1)

    for i,_ := range m1 {
        go func(tag string,result *[]model.GroundInfosID,lock sync.Mutex) {
            var result1 []model.GroundInfosID
            sql := "select name,user_picture,review_id,title,content,time,tag,picture,comment_sum,like_sum from user_review where tag = "+tag
            err := database.DB.Raw(sql).Scan(&result1).Error
            lock.Lock()
            (*result) = append((*result),result1...)
            lock.Unlock()
            errChannel <-err
        }(i,result,lock)
    }

    for i:= 0;i < l2;i++{
        if <-errChannel != nil {
            return <-errChannel
        }
    }

    return nil
}


func doReturnSearch(words []string,resultTagp *[]tag) error {
    var sum int
    var m1 map[string]bool
    l := len(words)
    //上面两个通道固定是接收l个信号
    channelCount := make(chan int,l)
    errChannel := make(chan error,l)
    channelIndex := make(chan string,2*l)
    defer close(channelCount)
    defer close(errChannel)
    defer close(channelIndex)

    for i := 0;i < l;i++{
        go func() {
            var tags []tag
            sql := "select tag from user_review where tag like '%"+words[i]+"%'"
            err := database.DB.Raw(sql).Scan(&tags).Error
            l := len(tags)
            errChannel <- err
            channelCount <- l
            for i := 0;i < l;i++{
                channelIndex <- tags[i].Tag
            }
        }()
    }

    for i := 0;i < l;i++{
        if <-errChannel != nil {
            return <-errChannel
        }
        sum += <-channelCount
    }

    for i := 0;i < sum ;i++{
        m1[<-channelIndex] = true
    }

    for i , _ := range m1 {
        tag1 := tag{
            Tag:i,
        }
        (*resultTagp) = append((*resultTagp),tag1)
    }

    return nil
}

