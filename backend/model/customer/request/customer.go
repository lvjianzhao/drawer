package request

import (
	"server/model/common/request"
)

type Customer struct {
	ID               uint   `json:"id"`                       // 主键ID
	Name             string `json:"name" validate:"required"` // 客户名称
	Type             uint   `json:"type" validate:"required"` // 客户类型
	BusinessName     string `json:"businessName"`             // 商务名称
	ServiceStatus    uint   `json:"serviceStatus"`            // 服务状态
	Description      string `json:"description"`              // 描述信息
	ServiceStartDate string `json:"serviceStartDate"`         // 服务起始时间
	ServiceEndDate   string `json:"serviceEndDate"`           // 服务截止时间
	MaintenanceType  int8   `json:"maintenanceType"`          // 维保类型(0:未知; 1:付费;2:免费)
}

// 客户名称分页条件查询及排序结构体
type CustomerSearchParams struct {
	IP string `json:"ip"`
	Customer
	request.PageInfo
}

// 客户事件搜索条件
type EventSearchParams struct {
	CustomerName string `json:"customerName"`
	EventKeyword string `json:"eventKeyword"`
	CustomerID   int    `json:"customerID"`
	Sort         string `json:"sort"`      // 排序字段
	StartTime    int64  `json:"startTime"` // 事件起始时间
	EndTime      int64  `json:"endTime"`   // 事件截止时间
	request.PageInfo
}
