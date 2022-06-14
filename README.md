# douyin
字节青训营抖音项目
### 用户登录
### 视频获取
### 视频列表
### 视频发布 
### 评论
### 关注
# 接口如下
### apiRouter.GET("/feed/", FeedVideos)
### apiRouter.GET("/user/", token.JwtMiddleware(), UserInfo)
### apiRouter.POST("/user/register/", UserRegister)
### apiRouter.POST("/user/login/", UserLogin)
### apiRouter.POST("/publish/action/", token.JwtMiddleware(), Publish)
### apiRouter.GET("/publish/list/", token.JwtMiddleware(), PublishList)