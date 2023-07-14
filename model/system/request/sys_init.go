package request

import (
	"demo-gin-1/config"
	"fmt"
)

type InitDB struct {
	DBType   string `json:"dbType"`                      // 数据库类型
	Host     string `json:"host"`                        // 服务器地址
	Port     string `json:"port"`                        // 端口
	UserName string `json:"userName" binding:"required"` // 用户名
	Password string `json:"password"`                    // 密码
	DBName   string `json:"dbName" binding:"required"`   // 数据库名
}

func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprint("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}

func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		GeneralDB: config.GeneralDB{
			Path:         i.Host,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "charset=uft8mb4&parseTime=True&loc=Local",
		},
	}
}
