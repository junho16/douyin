package main

import (
	"douyin/src/controller"
	"douyin/src/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	//r := gin.Default()
	//dao.SqlSession.AutoMigrate(&model.User{})
	//dao.SqlSession.AutoMigrate(&model.Video{})
	//dao.SqlSession.AutoMigrate(&model.Comment{})
	//dao.SqlSession.AutoMigrate(&model.Favorite{})
	//dao.SqlSession.AutoMigrate(&model.Following{})
	//dao.SqlSession.AutoMigrate(&model.Followers{})

	r := gin.Default()
	controller.InitRouter(r)

	r.Static("/data/videos", "data/videos")

	//注册路由
	//r := controller.InitRouter()
	//启动端口为8080的项目
	errRun := r.Run(":8080")
	if errRun != nil {
		return
	}

}
