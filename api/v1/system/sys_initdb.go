package system

import (
	"demo-gin-1/global"
	"demo-gin-1/model/common/response"
	"demo-gin-1/model/system/request"
	"demo-gin-1/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DBApi struct {
}

// InitDB 初始化用户数据库
func (i DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("已存在数据库配置")
		response.FailWithMessage("已存在数据库配置", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数校验不通过", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := system.InitDBService.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("自动创建数据库失败", zap.Error(err))
		response.FailWithMessage("自动创建数据库失败", c)
		return
	}
	response.OkWithMessage("自动创建数据库成功", c)
}
