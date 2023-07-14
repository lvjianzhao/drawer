package customer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	commReq "server/model/common/request"
	"server/model/common/response"
	customerReq "server/model/customer/request"
	"server/tools"
)

// GetCustomerEnv 获取客户环境信息
func (ca *CustomerApi) GetCustomerEnv(c *gin.Context) {
	var EnvSp customerReq.CustomerSearchParams
	EnvSp.Page, _ = tools.StringToInt(c.Query("currentPage"))
	EnvSp.PageSize, _ = tools.StringToInt(c.Query("size"))
	EnvSp.Name = c.Query("name")
	EnvSp.IP = c.Query("ip")
	EnvSp.ID, _ = tools.StringToUint(c.Query("customerId"))

	if list, total, err := customerService.GetCustomerEnv(EnvSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取客户列表失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			Page:     EnvSp.Page,
			PageSize: EnvSp.PageSize,
			Total:    total,
			List:     list,
		}, "获取成功", c)
	}
}

// AddOrUpdateCustomerEnv 添加/更新客户环境信息
func (ca *CustomerApi) AddOrUpdateCustomerEnv(c *gin.Context) {
	var customerEnv customerReq.CustomerEnv
	_ = c.ShouldBindJSON(&customerEnv)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&customerEnv); err != nil {
		msg := fmt.Sprintf("请求参数错误,error: %s", err)
		response.FailWithMessage(msg, c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if res, action, err := customerService.AddOrUpdateCustomerEnv(customerEnv); err != nil {
		msg := fmt.Sprintf("%s失败, %s", action, err)
		response.FailWithMessage(msg, c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		msg := fmt.Sprintf("%s成功", action)
		response.OkWithDetailed(res, msg, c)
	}
}

// DeleteCustomerEnv 删除客户环境信息
func (ca *CustomerApi) DeleteCustomerEnv(c *gin.Context) {
	id, _ := tools.StringToInt(c.Param("envId"))
	var cId commReq.CId
	_ = c.ShouldBindJSON(&cId)

	if err := customerService.DeleteCustomerEnv(id); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
