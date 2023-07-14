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

type CustomerApi struct{}

func (ca *CustomerApi) GetAllCustomer(c *gin.Context) {
	if list, total, err := customerService.GetAllCustomer(); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取客户列表失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			Total: total,
			List:  list,
		}, "获取成功", c)
	}
}

// GetCustomer 获取客户名称
func (ca *CustomerApi) GetCustomer(c *gin.Context) {
	var CustomerSp customerReq.CustomerSearchParams
	CustomerSp.Page, _ = tools.StringToInt(c.Query("currentPage"))
	CustomerSp.PageSize, _ = tools.StringToInt(c.Query("size"))
	CustomerSp.Name = c.Query("name")
	CustomerSp.ServiceStatus, _ = tools.StringToUint(c.Query("serviceStatus"))
	CustomerSp.MaintenanceType, _ = tools.StringToInt8(c.Query("maintenanceType"))
	CustomerSp.Type, _ = tools.StringToUint(c.Query("type"))

	if list, total, err := customerService.GetCustomer(CustomerSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取客户列表失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			Page:     CustomerSp.Page,
			PageSize: CustomerSp.PageSize,
			Total:    total,
			List:     list,
		}, "获取成功", c)
	}
}

// AddOrUpdateCustomer 添加/更新客户名称
func (ca *CustomerApi) AddOrUpdateCustomer(c *gin.Context) {
	var customer customerReq.Customer
	_ = c.ShouldBindJSON(&customer)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&customer); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if res, action, err := customerService.AddOrUpdateCustomer(customer); err != nil {
		msg := fmt.Sprintf("%s失败, %s", action, err)
		response.FailWithMessage(msg, c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		msg := fmt.Sprintf("%s成功", action)
		response.OkWithDetailed(res, msg, c)
	}
}

// DeleteCustomer 删除客户名称
func (ca *CustomerApi) DeleteCustomer(c *gin.Context) {
	id, _ := tools.StringToInt(c.Param("customerId"))
	var cId commReq.CId
	_ = c.ShouldBindJSON(&cId)

	if err := customerService.DeleteCustomer(id); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
