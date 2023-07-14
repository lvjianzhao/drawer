package response

import "time"

type LicenseResult struct {
	CreatedAt            time.Time `json:"createdAt"`
	ID                   uint      `json:"id"`
	License              string    `json:"license"`
	Name                 string    `json:"customerName"`
	Env                  string    `json:"env"`
	ExpiredDate          string    `json:"expiredDate"`
	ManagerApis          uint      `json:"managerApis"`
	ManagerUsers         uint      `json:"managerUsers"`
	ManagerProjects      uint      `json:"managerProjects"`
	ManagerIntegrations  uint      `json:"managerIntegrations"`
	ManagerOrganizations uint      `json:"managerOrganizations"`
	ManagerTests         uint      `json:"managerTests"`
	ManagerProducts      uint      `json:"managerProducts"`
	StudioUsers          uint      `json:"studioUsers"`
	StudioCompBlackList  string    `json:"studioCompBlackList"`
	PermissionMenu       string    `json:"permissionMenu"`
	CheckList            string    `json:"checkList"`
	Active               uint      `json:"active"`
}
