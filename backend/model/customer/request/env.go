package request

type CustomerEnv struct {
	ID           uint   `json:"id"`                             // 主键ID
	CustomerId   uint   `json:"customerId" validate:"required"` // 客户名称
	Env          string `json:"env"  validate:"required"`       // 环境
	Vpn          string `json:"vpn"  validate:"required"`       // VPN信息
	DeployReport string `json:"deploy_report"`                  // 部署报告
}
