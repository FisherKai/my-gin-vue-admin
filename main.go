package main

import (
	"demo-gin-1/router"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {
	r = gin.Default()
	//监听端口默认为8080
	router.InitRoute()
	r.Run(":8999")
}
