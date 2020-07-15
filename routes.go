package main

import (
	"github.com/gin-gonic/gin"
	"nano/controller"
	"nano/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine  {

	//此处控制用户接口
	userGroup := r.Group("/api/auth")
	{
		userGroup.POST("/login", controller.Login)
		userGroup.POST("/register", controller.Register)
		userGroup.POST("/info", middleware.AuthMiddleware() ,controller.Info)
		userGroup.POST("/postWeb", controller.PostWebsite)
		userGroup.POST("/getWeb", controller.GetUserWebsite)
	}
	//此处控制页面显示
	r.GET("/user/tool",controller.ShowWeblist)
	r.GET("/index",controller.ShowIndex)
	r.GET("/login",controller.ShowLogin)
	r.GET("/user",controller.Showuser)
	r.GET("/api/index/show", controller.ShowWebsite)

	//此处小工具的接口
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.Getdolist)
		v1Group.GET("/todo", controller.Showlist)
		v1Group.DELETE("/todo/:id",controller.Dellist)
		v1Group.PUT("/todo/:id",  controller.Doinglist)
	}

	return r
}