package response

import "time"

type CustomerResult struct {
	UpdatedAt        time.Time `json:"updatedAt"`
	ID               uint      `json:"id"`
	Name             string    `json:"name"`             // 客户名称
	Type             uint      `json:"type"`             // 客户类型
	BusinessName     string    `json:"businessName"`     // 商务名称
	ServiceStartDate time.Time `json:"serviceStartDate"` // 维保起始日期
	ServiceEndDate   time.Time `json:"serviceEndDate"`   // 维保截止日期
	ServiceStatus    uint      `json:"serviceStatus"`    // 维保状态
	Description      string    `json:"description"`      // 备注
	MaintenanceType  int8      `json:"maintenanceType"`  // 维保类型(0:未知; 1:付费;2:免费)
}
