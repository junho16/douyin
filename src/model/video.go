package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	AuthorId      uint   `json:"author"`         //视频作者信息
	PlayUrl       string `json:"play_url"`       //视频播放地址
	CoverUrl      string `json:"cover_url"`      //视频封面地址
	FavoriteCount uint   `json:"favorite_count"` //视频点赞总数
	CommentCount  uint   `json:"comment_count"`  //视频评论总数
	Title         string `json:"title"`          //视频标题
}
