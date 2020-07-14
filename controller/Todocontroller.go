package controller

import (
	"github.com/gin-gonic/gin"
	"nano/data"
	"nano/model"
	"net/http"
)

func Getdolist(c *gin.Context)  {
		DB :=data.GetDB()
		var todolist model.Todo
		c.BindJSON(&todolist)
		if err :=DB.Create(&todolist);err!=nil{
			c.JSON(200,gin.H{"err":err.Error})
		} else{
			c.JSON(200,todolist)
		}

}
func Showlist(c *gin.Context)  {
	DB :=data.GetDB()
	var todoList []model.Todo
	if err := DB.Find(&todoList).Error; err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, todoList)
	}
}
func Dellist(c *gin.Context)  {
	DB :=data.GetDB()
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := DB.Where("id=?", id).Delete(model.Todo{}).Error;err!=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}