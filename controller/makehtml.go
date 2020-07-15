package controller

import (
	"github.com/gin-gonic/gin"
	"nano/data"
	"nano/model"
	"net/http"
)
type webInfo struct {
	Website string
	Num    int
}

func ShowIndex(c *gin.Context)  {
	DB := data.GetDB()
	var web []model.Allweb

	DB.Find(&web)
	//fmt.Printf("%s\n",web[0].Website)
	if len(web)==0 {
		c.JSON(402,gin.H{"msg":"没有收藏过一个网页"})
		return
	}
	var aweb []webInfo
	for  i:=0 ;i<len(web);i++ {
		aweb= append(aweb, webInfo{web[i].Website,web[i].Num})

	}


	c.HTML(http.StatusOK,"index.html",aweb)
}
func ShowLogin(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",nil)
}
func Showuser(c *gin.Context) {
	c.HTML(http.StatusOK,"share.html",nil)
}
func ShowWeblist(c *gin.Context)  {
	c.HTML(http.StatusOK,"checklist.html",nil)
}
