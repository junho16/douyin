package controller

import (
	"douyin/src/common"
	"douyin/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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

func GetFeedVideos(c *gin.Context) {
	strToken := c.Query("token")
	strLastTime := c.Query("latest_time")
	fmt.Println(strToken, " ", strLastTime)
	//var haveToken bool
	//if strToken == "" {
	//	haveToken = false
	//} else {
	//	haveToken = true
	//}

	lastTime, err := strconv.ParseInt(strLastTime, 10, 32)
	if err != nil {
		lastTime = 0
	}

	fmt.Println(lastTime)
	//var feedVideoList []FeedVideo
	//feedVideoList = make([]FeedVideo, 0)
	videoList, _ := service.FeedGet(lastTime)
	//var newTime int64 = 0 //返回的视频的最久的一个的时间

	for _, x := range videoList {
		fmt.Println(x)
	}

}
