package internal

import (
	"demo-gin-1/global"
	"fmt"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{
		Writer: w,
	}
}

// 格式化日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GVA_CONFIG.Mysql.LogZap
	}
	if logZap {
		global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
