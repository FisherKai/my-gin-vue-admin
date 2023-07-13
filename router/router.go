package router

import (
	"demo-gin-1/db"
	"fmt"
)

func InitRoute() {
	db.InitDB()
	fmt.Println("初始化路由......")
	// Get
	getUsernameAndAction()
	fmt.Println("初始化路由✅......")
}
