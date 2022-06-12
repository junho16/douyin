package controller

import (
	"douyin/src/common"
	"douyin/src/service"
	"douyin/src/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedVideo struct {
	Id            uint     `json:"id,omitempty"`
	Author        FeedUser `json:"author,omitempty"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount uint     `json:"favorite_count,omitempty"`
	CommentCount  uint     `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         string   `json:"title,omitempty"`
}

type FeedUser struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   uint   `json:"follow_count,omitempty"`
	FollowerCount uint   `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type DouyinFeedResponse struct {
	common.Response
	VideoList []FeedVideo `json:"video_list,omitempty"`
	NextTime  uint        `json:"next_time,omitempty"`
}
type DouyinFeedNoVideoResponse struct {
	common.Response
	NextTime uint `json:"next_time"`
}

func FeedVideos(c *gin.Context) {
	strToken := c.Query("token")
	strLastTime := c.Query("latest_time")
	fmt.Println(strToken, " ", strLastTime)

	//判断token是否有效
	var tokenValid bool
	tokenStruct, ok := token.CheckToken(strToken)
	if ok && time.Now().Unix() <= tokenStruct.ExpiresAt {
		tokenValid = false
	} else {
		tokenValid = true
	}

	lastTime, err := strconv.ParseInt(strLastTime, 10, 32)
	if err != nil {
		lastTime = 0
	}

	fmt.Println(lastTime)
	feedVideoList := make([]FeedVideo, 0)
	videoList, _ := service.FeedGet(lastTime)
	var newTime int64 = 0 //返回的视频的最久的一个的时间

	for _, x := range videoList {
		var feedTmp FeedVideo
		feedTmp.Id = x.ID
		feedTmp.PlayUrl = x.PlayUrl
		// Author 需要从用户信息接口中查询
		var user, err = service.GetUser(x.AuthorId)
		var feedUser FeedUser
		if err == nil { //用户存在
			feedUser.Id = user.ID
			feedUser.FollowerCount = user.FollowerCount
			feedUser.FollowCount = user.FollowCount
			feedUser.Name = user.Name
			feedUser.IsFollow = false
			if tokenValid {
				// 查询是否关注
				var uid1 = tokenStruct.UserId //用户id
				var uid2 = x.AuthorId         //视频发布者id
				if service.IsFollowing(uid1, uid2) {
					//user2是否已经被user1关注
					feedUser.IsFollow = true
				}
			}
		}
		feedTmp.Author = feedUser
		feedTmp.CommentCount = x.CommentCount
		feedTmp.CoverUrl = x.CoverUrl
		feedTmp.FavoriteCount = x.FavoriteCount
		feedTmp.IsFavorite = false
		if tokenValid {
			//查询是否点赞过
			var uid = tokenStruct.UserId         //用户id
			var vid = x.ID                       // 视频id
			if service.CheckFavorite(uid, vid) { //有点赞记录
				feedTmp.IsFavorite = true
			}
		}
		feedTmp.Title = x.Title
		feedVideoList = append(feedVideoList, feedTmp)
		newTime = x.CreatedAt.Unix()
	}
	if len(feedVideoList) > 0 {
		c.JSON(http.StatusOK, DouyinFeedResponse{
			Response:  common.Response{StatusCode: 0}, //成功
			VideoList: feedVideoList,
			NextTime:  uint(newTime),
		})
	} else {
		c.JSON(http.StatusOK, DouyinFeedNoVideoResponse{
			Response: common.Response{StatusCode: 0}, //成功
			NextTime: 0,                              //重新循环
		})
	}
}
