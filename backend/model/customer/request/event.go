package request

type CustomerEvent struct {
	ID           uint   `json:"id"`                                // 主键ID
	CustomerId   uint   `json:"customerID" validate:"required"`    // 客户名称
	EventTime    string `json:"eventTime" validate:"required"`     // 事件时间
	EventContent string `json:"eventContent"  validate:"required"` // 事件内容
}
