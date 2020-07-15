package main

import (
	"github.com/gin-gonic/gin"
	"nano/data"
)

func main() {
	//初始化数据库
	db :=data.InitDB()
	defer db.Close()

	r:= gin.Default()

	// 指定HTML以及渲染页面
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")

	//启动
	r=CollectRoute(r)
	
	r.Run()
}

