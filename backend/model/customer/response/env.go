package response

import "time"

type CustomerEnvResult struct {
	UpdatedAt    time.Time `json:"updatedAt"`
	ID           uint      `json:"id"`                                 // 主键ID
	Name         string    `json:"name" validate:"required"`           // 客户名称
	CustomerId   uint      `json:"customerId" validate:"required"`     // 客户ID
	Env          string    `json:"env"  validate:"required"`           // 环境
	Vpn          string    `json:"vpn"  validate:"required"`           // VPN信息
	DeployReport string    `json:"deploy_report"  validate:"required"` // 部署报告
}
