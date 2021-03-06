package controller

import (
	"douyin/src/common"
	"douyin/src/model"
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

type VideoListResponse struct {
	common.Response
	VideoList []ReturnVideo `json:"video_list"`
}

type ReturnAuthor struct {
	AuthorId      uint   `json:"author_id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type ReturnVideo struct {
	VideoId       uint         `json:"video_id"`
	Author        ReturnAuthor `json:"author"`
	PlayUrl       string       `json:"play_url"`
	CoverUrl      string       `json:"cover_url"`
	FavoriteCount uint         `json:"favorite_count"`
	CommentCount  uint         `json:"comment_count"`
	IsFavorite    bool         `json:"is_favorite"`
	Title         string       `json:"title"`
}

func FeedVideos(c *gin.Context) {
	strToken := c.Query("token")
	strLastTime := c.Query("latest_time")
	fmt.Println(strToken, " ", strLastTime)

	//??????token????????????
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
	var newTime int64 = 0 //??????????????????????????????????????????

	for _, x := range videoList {
		var feedTmp FeedVideo
		feedTmp.Id = x.ID
		feedTmp.PlayUrl = x.PlayUrl
		// Author ????????????????????????????????????
		var user, err = service.GetUser(x.AuthorId)
		var feedUser FeedUser
		if err == nil { //????????????
			feedUser.Id = user.ID
			feedUser.FollowerCount = user.FollowerCount
			feedUser.FollowCount = user.FollowCount
			feedUser.Name = user.Name
			feedUser.IsFollow = false
			if tokenValid {
				// ??????????????????
				var uid1 = tokenStruct.UserId //??????id
				var uid2 = x.AuthorId         //???????????????id
				if service.IsFollowing(uid1, uid2) {
					//user2???????????????user1??????
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
			//?????????????????????
			var uid = tokenStruct.UserId         //??????id
			var vid = x.ID                       // ??????id
			if service.CheckFavorite(uid, vid) { //???????????????
				feedTmp.IsFavorite = true
			}
		}
		feedTmp.Title = x.Title
		feedVideoList = append(feedVideoList, feedTmp)
		newTime = x.CreatedAt.Unix()
	}
	if len(feedVideoList) > 0 {
		c.JSON(http.StatusOK, DouyinFeedResponse{
			Response:  common.Response{StatusCode: 0}, //??????
			VideoList: feedVideoList,
			NextTime:  uint(newTime),
		})
	} else {
		c.JSON(http.StatusOK, DouyinFeedNoVideoResponse{
			Response: common.Response{StatusCode: 0}, //??????
			NextTime: 0,                              //????????????
		})
	}
}

func PublishList(c *gin.Context) { //?????????????????????
	//1.???????????????token
	getHostId, _ := c.Get("user_id")
	var HostId uint
	if v, ok := getHostId.(uint); ok {
		HostId = v
	}
	//2.????????????????????????id??????????????????????????????
	getGuestId := c.Query("user_id")
	id, _ := strconv.Atoi(getGuestId)
	GuestId := uint(id)

	//????????????id????????????
	getUser, err := service.GetUser(GuestId)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "Not find this person.",
		})
		c.Abort()
		return
	}

	returnAuthor := ReturnAuthor{
		AuthorId:      GuestId,
		Name:          getUser.Name,
		FollowCount:   getUser.FollowCount,
		FollowerCount: getUser.FollowerCount,
		IsFollow:      service.IsFollowing(HostId, GuestId),
	}
	//????????????id?????? ????????????????????????
	var videoList []model.Video
	videoList = service.GetVideoList(GuestId)
	if len(videoList) == 0 {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "null",
			},
			VideoList: nil,
		})
	} else { //???????????????????????????
		var returnVideoList []ReturnVideo
		for i := 0; i < len(videoList); i++ {
			returnVideo := ReturnVideo{
				VideoId:       videoList[i].ID,
				Author:        returnAuthor,
				PlayUrl:       videoList[i].PlayUrl,
				CoverUrl:      videoList[i].CoverUrl,
				FavoriteCount: videoList[i].FavoriteCount,
				CommentCount:  videoList[i].CommentCount,
				IsFavorite:    service.CheckFavorite(HostId, videoList[i].ID),
				Title:         videoList[i].Title,
			}
			returnVideoList = append(returnVideoList, returnVideo)
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			VideoList: returnVideoList,
		})
	}
}
