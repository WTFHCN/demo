package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nano/data"
	"nano/model"
	"net/http"
	"sort"
)
type webInfo struct {
	Website string
	Image string
	Num    int
}

func min(a int, b int ) int   {
	if a<b	{
		return a
	}else {
		return b
	}
}
func ShowIndex(c *gin.Context)  {
	DB := data.GetDB()
	var web []model.Allweb

	DB.Find(&web)
	//fmt.Printf("%s\n",web[0].Website)


	var aweb []webInfo

	for  i:=0 ;i<min(6,len(web));i++ {
		if web[i].Num>0 {
			var o model.Webimage
			fmt.Printf("%s\n",web[i].Website)
			DB.Debug().Where("website = ?",web[i].Website).First(&o)

			if o.ID!=0 {
				aweb = append(aweb, webInfo{web[i].Website,"static/images/"+o.Image, web[i].Num})
			}
		}
	}
	less := func(i, j int) bool {
		return aweb[i].Num > aweb[j].Num
	}
	sort.Slice(aweb,less)

	c.HTML(http.StatusOK,"index.html",aweb)
}
func ShowLogin(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",nil)
}
func Showuser(c *gin.Context) {
	c.HTML(http.StatusOK,"user.html",nil)
}
func ShowWeblist(c *gin.Context)  {
	c.HTML(http.StatusOK,"checklist.html",nil)
}
