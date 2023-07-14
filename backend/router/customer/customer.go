package customer

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type CustomerRouter struct{}

func (u *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	customerRouter := Router.Group("customer")
	customerApi := api.ApiGroupApp.CustomerGroup.CustomerApi
	{
		// 客户维保信息
		customerRouter.GET("", customerApi.GetCustomer)
		customerRouter.POST("", customerApi.AddOrUpdateCustomer)
		customerRouter.PUT("", customerApi.AddOrUpdateCustomer)
		customerRouter.DELETE("/:customerId", customerApi.DeleteCustomer)
		customerRouter.GET("/all", customerApi.GetAllCustomer)
		// 客户环境
		customerRouter.GET("/env", customerApi.GetCustomerEnv)
		customerRouter.GET("/env/:customerId", customerApi.GetCustomerEnv)
		customerRouter.POST("/env", customerApi.AddOrUpdateCustomerEnv)
		customerRouter.PUT("/env", customerApi.AddOrUpdateCustomerEnv)
		customerRouter.DELETE("/env/:envId", customerApi.DeleteCustomerEnv)
		// 客户事件
		customerRouter.GET("/event", customerApi.GetCustomerEvent)
		customerRouter.POST("/event", customerApi.AddOrUpdateEvent)
		customerRouter.PUT("/event", customerApi.AddOrUpdateEvent)
		customerRouter.DELETE("/event/:eventId", customerApi.DeleteEvent)

		// 客户数据分析
		customerRouter.GET("/summary", customerApi.GetCustomerSummary)
	}
	return customerRouter
}
