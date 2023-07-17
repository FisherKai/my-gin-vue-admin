package initialize

import (
	"demo-gin-1/config"
	"demo-gin-1/global"
	"demo-gin-1/initialize/internal"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化mysql数据库
func GormMysql() *gorm.DB {
	fmt.Println("mysql 数据库准备开始初始化......")
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, //根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		return db
	}
}

func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		return db
	}
}
