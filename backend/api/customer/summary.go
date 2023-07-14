package customer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
)

// GetCustomerSummary  获取客户分析所需的数据，对应dashboard页面
func (ca *CustomerApi) GetCustomerSummary(c *gin.Context) {
	if list, err := customerService.GetCustomerSummary(); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取客户列表失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
