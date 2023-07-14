package license

import (
	"server/global"
	"server/model/customer"
	"time"
)

type LicenseModel struct {
	global.TD27_MODEL
	CustomerID           uint                   `json:"customerID" gorm:"not null;comment:客户ID;index" validate:"required"`
	Customer             customer.CustomerModel `gorm:"foreignKey:CustomerID"`
	Env                  string                 `json:"env" gorm:"not null;comment:环境" validate:"required"`
	ExpiredDate          time.Time              `json:"expiredDate" gorm:"not null;comment:license有效期" validate:"required"`
	License              string                 `json:"license" gorm:"type:text;comment:license"`
	ManagerApis          uint                   `json:"managerApis" gorm:"not null;comment:管理平台API数"`
	ManagerUsers         uint                   `json:"managerUsers" gorm:"not null;comment:管理平台用户数"`
	ManagerProjects      uint                   `json:"managerProjects" gorm:"not null;comment:管理平台项目数"`
	ManagerIntegrations  uint                   `json:"managerIntegrations" gorm:"not null;comment:管理平台集成平台限制"`
	ManagerOrganizations uint                   `json:"managerOrganizations" gorm:"not null;comment:管理平台支持组织数"`
	ManagerTests         uint                   `json:"managerTests" gorm:"not null;comment:管理平台测试限制"`
	ManagerProducts      uint                   `json:"managerProducts" gorm:"not null;comment:管理平台发布限制"`
	StudioUsers          uint                   `json:"studioUsers" gorm:"not null;comment:集成平台用户数"`
	StudioCompBlackList  string                 `json:"studioCompBlackList" gorm:"type:text;comment:集成组件黑名单"`
	PermissionMenu       string                 `json:"permissionMenu" gorm:"type:text;not null;comment:权限设置"`
	CheckList            string                 `json:"checkList" gorm:"type:text;not null;comment:权限设置复选框"`
	Active               uint                   `json:"active" gorm:"default:1;comment:license是否生效(1:生效中,2:已作废)"`
}

func (LicenseModel) TableName() string {
	return "license"
}
