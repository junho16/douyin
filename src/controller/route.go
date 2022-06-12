package controller

import (
	"douyin/src/token"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	//r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	//apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", token.JwtMiddleware(), UserInfo)
	apiRouter.POST("/user/register/", UserRegister)
	apiRouter.POST("/user/login/", UserLogin)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)
	//
	//// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}

//func InitRouter() *gin.Engine {
//	r := gin.Default()
//	// 主路由组
//	douyinGroup := r.Group("/douyin")
//	{
//		// user路由组
//		userGroup := douyinGroup.Group("/user")
//		{
//			userGroup.GET("/", token.JwtMiddleware(), UserInfo)
//			userGroup.POST("/login/", token.JwtMiddleware(), UserLogin)
//			userGroup.POST("/register/", token.JwtMiddleware(), UserRegister)
//		}
//		//
//		//// publish路由组
//		//publishGroup := douyinGroup.Group("/publish")
//		//{
//		//	publishGroup.POST("/action/", middleware.JwtMiddleware(), Publish)
//		//	publishGroup.GET("/list/", middleware.JwtMiddleware(), PublishList)
//		//
//		//}
//		//
//		//// feed
//		//douyinGroup.GET("/feed/", controller.Feed)
//		//
//		//favoriteGroup := douyinGroup.Group("favorite")
//		//{
//		//	favoriteGroup.POST("/action", middleware.JwtMiddleware(), controller.Favorite)
//		//	favoriteGroup.GET("/list", middleware.JwtMiddleware(), controller.FavoriteList)
//		//}
//		//
//		//// comment路由组
//		//commentGroup := douyinGroup.Group("/comment")
//		//{
//		//	commentGroup.POST("/action/", middleware.JwtMiddleware(), controller.CommentAction)
//		//	commentGroup.GET("/list/", middleware.JwtMiddleware(), controller.CommentList)
//		//}
//		//
//		//// relation路由组
//		//relationGroup := douyinGroup.Group("relation")
//		//{
//		//	relationGroup.POST("/action/", middleware.JwtMiddleware(), controller.RelationAction)
//		//	relationGroup.GET("/follow/list/", middleware.JwtMiddleware(), controller.FollowList)
//		//	relationGroup.GET("/follower/list/", middleware.JwtMiddleware(), controller.FollowerList)
//		//}
//	}
//
//	return r
//}
