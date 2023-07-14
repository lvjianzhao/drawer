package customer

import (
	"server/global"
	"time"
)

type CustomerEventModel struct {
	global.TD27_MODEL
	CustomerID   uint          `json:"customerID" gorm:"not null;comment:客户ID;index" validate:"required"`
	Customer     CustomerModel `gorm:"foreignKey:CustomerID"`
	EventTime    time.Time     `json:"eventTime" gorm:"not null;comment:事件发生日期;type:datetime"`
	EventContent string        `json:"eventContent" gorm:"not null;type:text;comment:事件内容"  validate:"required"`
}

func (CustomerEventModel) TableName() string {
	return "customer_event"
}
