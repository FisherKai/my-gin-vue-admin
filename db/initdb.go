package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB() {
	fmt.Println("初始化数据库......")
	var err error
	const dataSourceName = user + ":" + passwd + "@tcp(" + host + ":3306)/" + dbname
	const dbType = "mysql"
	Db, err = sql.Open(dbType, dataSourceName)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	fmt.Println("初始化数据库✅......")
}
