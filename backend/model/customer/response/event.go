package response

import "time"

type CustomerEventResult struct {
	UpdatedAt    time.Time `json:"updatedAt"`
	ID           uint      `json:"id"`                               // 主键ID
	Name         string    `json:"name" validate:"required"`         // 客户名称
	CustomerId   uint      `json:"customerID" validate:"required"`   // 客户ID
	EventTime    time.Time `json:"eventTime"`                        // 事件发生日期
	EventContent string    `json:"eventContent" validate:"required"` // 事件内容
}
