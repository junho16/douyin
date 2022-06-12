package service

import (
	"douyin/src/dao"
	"douyin/src/model"
	"fmt"
)

// CreateVideo 添加一条视频信息
func CreateVideo(video *model.Video) {
	//dao.SqlSession.Table("videos").Create(&video)

	if err := dao.SqlSession.Model(&model.Video{}).Create(&video).Error; err != nil {
		//错误处理
		fmt.Println(err)
	}
}
