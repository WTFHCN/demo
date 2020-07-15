package controller

import (
	"github.com/gin-gonic/gin"
	"nano/data"
	"nano/model"
	"net/http"
)

func PostWebsite(c *gin.Context)  {
	DB := data.GetDB()
	name := c.PostForm("name")
	website := c.PostForm("website")
	introduction := c.PostForm("introduction")
	if website == " " ||introduction == " "{
		c.JSON(http.StatusOK,gin.H{"msg":"输入格式有问题"})
		return
	}

	u :=model.Webcount{
		Name: name,
		Website: website,
		Introduction: introduction,

	}
	DB.Create(&u)
	var o model.Allweb
	DB.Where("website = ?",website).First(&o)

	if o.ID == 0{
		DB.Create(&model.Allweb{Website: website,Num: 1})
		c.JSON(http.StatusOK,gin.H{"msg":"添加成功，恭喜你是第一个添加网站的"})
		return
	}
	o.Num=o.Num+1
	DB.Save(&o)
	c.JSON(http.StatusOK,gin.H{"msg":"添加成功"})


}
func GetUserWebsite(c *gin.Context)  {
	DB := data.GetDB()
	var web []model.Webcount
	name := c.PostForm("name")
	DB.Where("name = ?", name).Find(&web)


	c.JSON(http.StatusOK,web)

}
func ShowWebsite(c *gin.Context)  {
	DB := data.GetDB()
	var web []model.Allweb

	DB.Find(&web)

	c.JSON(http.StatusOK,web)


}
func AddWeb(c*gin.Context)  {
	DB:=data.GetDB()
	website:= c.PostForm("website")
	var o model.Allweb
	DB.Debug().Where("website = ?",website).First(&o)
	o.Num=o.Num+1
	//fmt.Printf("%d\n",o.Num)
	DB.Save(&o)
	c.JSON(http.StatusOK,gin.H{"msg":"点赞成功"})
}