package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var r = gin.Default()

func getUsernameAndAction() {
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
}
