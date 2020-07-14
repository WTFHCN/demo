package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowIndex(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}
func ShowLogin(c *gin.Context)  {
	c.HTML(http.StatusOK,"login.html",nil)
}