package main

import (
	"github.com/gin-gonic/gin"
	"nano/controller"
	"nano/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine  {

	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/info", middleware.AuthMiddleware() ,controller.Info)
	r.GET("/index",controller.ShowIndex)
	r.GET("/login",controller.ShowLogin)
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.Getdolist)
		v1Group.GET("/todo", controller.Showlist)
		v1Group.DELETE("/todo/:id",controller.Dellist)
	}

	return r
}