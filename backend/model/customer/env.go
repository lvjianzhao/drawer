package customer

import (
	"server/global"
)

type CustomerEnvModel struct {
	global.TD27_MODEL
	CustomerID   uint          `json:"customerId" gorm:"not null;uniqueIndex:idx_customerId_env;comment:客户ID;index" validate:"required"`
	Customer     CustomerModel `gorm:"foreignKey:CustomerID"`
	Env          string        `json:"env" gorm:"not null;uniqueIndex:idx_customerId_env;comment:环境" validate:"required"`
	Vpn          string        `json:"vpn" gorm:"not null;type:text;comment:VPN信息"  validate:"required"`
	DeployReport string        `json:"deploy_report" gorm:"not null;type:text;comment:部署报告"  validate:"required"`
}

func (CustomerEnvModel) TableName() string {
	return "customer_env"
}
