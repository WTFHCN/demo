package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"nano/data"

	"nano/model"
	"net/http"

)

func Register(ctx *gin.Context) {
	DB :=data.GetDB()

	name := ctx.PostForm("name")
	password := ctx.PostForm("password") // 可设置默认值
	if len(password) < 6 {
		ctx.JSON(422, gin.H{
			"code": 422,
			"msh":  "密码太短",
		})
		return
	}


	newUser := model.User{
		Name:     name,
		Password: password,
	}
	log.Printf("2333")
	DB.Create(&newUser)


	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"password": password,
	})
}

func Login (ctx *gin.Context) {
	DB :=data.GetDB()
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	var user model.User
	DB.Where("name = ?", name).First(&user)
	if user.Password != password {
		ctx.JSON(402, gin.H{
			"msg":"登陆失败",
			"name": name,
			"password": password,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":"登陆成功",
		"name": name,
		"password": password,
	})
}

