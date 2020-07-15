package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"nano/data"
	"nano/model"
	"net/http"
)

func Register(ctx *gin.Context) {
	DB :=data.GetDB()
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	fmt.Printf("%s %s \n",name,password)
	var user model.User
	DB.Where("name = ?", name).First(&user)

	if user.ID != 0 {
		ctx.JSON(200, gin.H{
			"msg":"用户已经存在",
		})
		return
	}
	if len(password) < 6 {

		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "密码太短",
		})
		return
	}
	// 创建用户

	newUser := model.User{
		Name:     name,
		Password: password,
	}

	DB.Create(&newUser)


	ctx.JSON(http.StatusOK, gin.H{
		"msg":"注册成功",
	})
}

func Login (ctx *gin.Context) {
	DB :=data.GetDB()
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	fmt.Printf("%s %s \n",name,password)
	var user model.User
	DB.Where("name = ?", name).First(&user)

	if user.ID == 0 {
		ctx.JSON(200, gin.H{
			"msg":"用户不存在",

		})

		return
	}
	if user.Password != password {
		ctx.JSON(200, gin.H{
			"msg":"登录失败",
			"name": name,

		})
		return
	}
	token ,err:= ReleaseToken(user)
	if err !=nil {
		ctx.JSON(200, gin.H{
			"msg":"生成token失败",
		})
		log.Printf("token error")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":"登录成功",
		"date":gin.H{"token":token},
	})
}
func Info(ctx *gin.Context)  {
	user, _:=ctx.Get("user")
	ctx.JSON(200, gin.H{
		"user":user,
	})
}
