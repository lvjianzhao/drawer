package customer

import (
	"gorm.io/gorm"
	"server/global"
	"time"
)

type CustomerModel struct {
	global.TD27_MODEL
	Name             string    `json:"name" gorm:"not null;unique;comment:客户名称" validate:"required"`
	Type             uint      `json:"type" gorm:"not null;default:1;comment:客户类型(1:POC 2:正式)"`
	BusinessName     string    `json:"businessName" gorm:"comment:商务名称" `
	ServiceStartDate time.Time `json:"serviceStartDate" gorm:"not null;comment:维保起始日期;type:date;default:2023-01-01""`
	ServiceEndDate   time.Time `json:"serviceEndDate" gorm:"not null;comment:维保起始日期;type:date;default:2023-01-01""`
	ServiceStatus    uint      `json:"serviceStatus" gorm:"not null;comment:维保状态: (0: 已过期,1: 维保中)"`
	MaintenanceType  int8      `json:"maintenanceType" gorm:"not null;default:0;comment:维保类型(0:未知; 1:付费;2:免费)"`
	Description      string    `json:"description" gorm:"type:text;comment:描述信息"`
}

func (c *CustomerModel) BeforeSave(tx *gorm.DB) error {
	c.ServiceStatus = 0
	if c.ServiceEndDate.IsZero() {
		return nil
	}
	// 检查 ServiceEndDate 是否比今天大
	if c.ServiceEndDate.After(time.Now().Local()) {
		c.ServiceStatus = 1
	}

	return nil
}

func (CustomerModel) TableName() string {
	return "customer"
}
