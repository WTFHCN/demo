package main

import (
	"github.com/gin-gonic/gin"
	"nano/data"
)



func main() {
	db :=data.InitDB()
	defer db.Close()
	gin.SetMode(gin.ReleaseMode)
	r:= gin.Default()
	r=CollectRoute(r)
	panic(r.Run())
}

