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

// AddOrUpdateEvent 添加/更新客户事件
func (ca *CustomerApi) AddOrUpdateEvent(c *gin.Context) {
	var event customerReq.CustomerEvent
	_ = c.ShouldBindJSON(&event)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&event); err != nil {
		msg := fmt.Sprintf("请求参数错误,error: %s", err)
		response.FailWithMessage(msg, c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if res, action, err := customerService.AddOrUpdateEvent(event); err != nil {
		msg := fmt.Sprintf("%s失败, %s", action, err)
		response.FailWithMessage(msg, c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		msg := fmt.Sprintf("%s成功", action)
		response.OkWithDetailed(res, msg, c)
	}
}

// GetCustomerEvent  获取客户事件
func (ca *CustomerApi) GetCustomerEvent(c *gin.Context) {
	var EventSp customerReq.EventSearchParams
	EventSp.Page, _ = tools.StringToInt(c.Query("currentPage"))
	EventSp.PageSize, _ = tools.StringToInt(c.Query("size"))
	EventSp.CustomerName = c.Query("customerName")
	EventSp.EventKeyword = c.Query("eventKeyword")
	EventSp.CustomerID, _ = tools.StringToInt(c.Query("customerID"))
	EventSp.Sort = c.Query("sort")
	EventSp.StartTime, _ = tools.StringToInt64(c.Query("startTime"))
	EventSp.EndTime, _ = tools.StringToInt64(c.Query("endTime"))

	if list, total, err := customerService.GetCustomerEvent(EventSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取客户列表失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			Page:     EventSp.Page,
			PageSize: EventSp.PageSize,
			Total:    total,
			List:     list,
		}, "获取成功", c)
	}
}

// DeleteEvent 删除客户事件
func (ca *CustomerApi) DeleteEvent(c *gin.Context) {
	id, _ := tools.StringToInt(c.Param("eventId"))
	var cId commReq.CId
	_ = c.ShouldBindJSON(&cId)

	if err := customerService.DeleteEvent(id); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
