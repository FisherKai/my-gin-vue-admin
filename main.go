package main

import (
	"demo-gin-1/core"
	"demo-gin-1/global"
	"demo-gin-1/initialize"
	"fmt"
)

func main() {
	global.GVA_VP = core.Viper()
	fmt.Println("系统准备开始初始化......")
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		// 初始化表
		initialize.RegisterTables()
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

}
