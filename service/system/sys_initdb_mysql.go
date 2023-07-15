package system

import (
	"context"
	"demo-gin-1/config"
	"demo-gin-1/global"
	"demo-gin-1/model/system/request"
	"demo-gin-1/utils"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path/filepath"
)

type MysqlInitHandler struct {
}

func (h MysqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	//TODO implement me
	return nil
}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// mysql配置回写
func (h MysqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mysql)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.GVA_CONFIG.System.DbType = "mysql"
	global.GVA_CONFIG.Mysql = c
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	return global.GVA_VP.WriteConfig()
}

// 创建数据库并初始化mysql
func (h MysqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, ErrDBTypeMismatch
	}
	c := conf.ToMysqlConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		fmt.Println("无数据库名称，即将退出初始化...")
		return ctx, nil
	}
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.Dbname)
	// 创建数据库
	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		return nil, err
	}

	var db *gorm.DB
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return ctx, err
	}
	// filepath.Abs 获取绝对路径
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
	return nil
}

func InitData(ctx context.Context, inits initSlice) error {
	//TODO implement me
	return nil
}
