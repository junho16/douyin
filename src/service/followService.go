package service

import (
	"douyin/src/dao"
	"douyin/src/model"
	"github.com/jinzhu/gorm"
)

// 判断HostId是否已经关注GuestId
func IsFollowing(HostId uint, GuestId uint) bool {
	var relationExist = &model.Following{}
	if err := dao.SqlSession.Model(&model.Following{}).Where("host_id=? AND guest_id=?", HostId, GuestId).First(&relationExist).Error; gorm.IsRecordNotFoundError(err) {
		//关注不存在
		return false
	}
	//关注存在
	return true
}
