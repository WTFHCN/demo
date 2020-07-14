package main

import (
	"github.com/gin-gonic/gin"
	"nano/data"
)

func main() {
	
	db :=data.InitDB()
	defer db.Close()
	r:= gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")
	r=CollectRoute(r)
	
	r.Run()
}

