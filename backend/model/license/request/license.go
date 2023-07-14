package request

import (
	"server/model/common/request"
)

type License struct {
	ID                   uint    `json:"id"`                             // 主键ID
	CustomerID           uint    `json:"customerID" validate:"required"` // 客户ID
	CustomerName         string  `json:"customerName"`
	ENV                  string  `json:"env" validate:"required"` // 客户类型
	CompDisplayName      string  `json:"compDisplayName" `        // 公司显示名
	CompName             string  `json:"compName" `               // 公司名
	CompFullName         string  `json:"compFullName" `           // 公司全名
	ExpiredDate          string  `json:"expiredDate"`             // 过期时间
	ManagerApis          uint    `json:"managerApis"`             // 管理平台API数
	ManagerUsers         uint    `json:"managerUsers"`            // 管理平台用户数
	ManagerProjects      uint    `json:"managerProjects"`         // 管理平台用户数
	ManagerIntegrations  uint    `json:"managerIntegrations"`     // 管理平台用户数
	ManagerOrganizations uint    `json:"managerOrganizations"`    // 管理平台支持组织数
	ManagerTests         uint    `json:"managerTests"`            // 管理平台测试限制
	ManagerProducts      uint    `json:"managerProducts"`         // 管理平台测试限制
	StudioUsers          uint    `json:"studioUsers"`             // 管理平台测试限制
	StudioCompBlackList  string  `json:"studioCompBlackList"`     // 集成组件黑名单
	PermissionMenu       string  `json:"permissionMenu"`          // 权限设置
	CheckList            string  `json:"checkList"`
	CompTel              *string `json:"compTel"` // 电话
}

// license分页条件查询及排序结构体
type LicenseSearchParams struct {
	License
	Active uint `json:"active"` // license状态
	request.PageInfo
}
