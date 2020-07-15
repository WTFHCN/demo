package main

import (
	"github.com/gin-gonic/gin"
	"nano/controller"
	"nano/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine  {

	userGroup := r.Group("/api/auth")
	{
		userGroup.POST("/login", controller.Login)
		userGroup.POST("/register", controller.Register)
		userGroup.POST("/info", middleware.AuthMiddleware() ,controller.Info)
		userGroup.POST("/postWeb", controller.PostWebsite)
		userGroup.POST("/getWeb", controller.GetUserWebsite)
	}

	r.GET("/index",controller.ShowIndex)
	r.GET("/login",controller.ShowLogin)

	r.GET("/api/index/show", controller.ShowWebsite)
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.Getdolist)
		v1Group.GET("/todo", controller.Showlist)
		v1Group.DELETE("/todo/:id",controller.Dellist)
	}

	return r
}