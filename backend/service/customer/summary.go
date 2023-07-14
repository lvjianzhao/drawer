package customer

import (
	"server/global"
	CustomerModel "server/model/customer"
	customerRes "server/model/customer/response"
	LicenseModel "server/model/license"
)

func getTypeName(t uint) string {
	if t == 2 {
		return "正式"
	} else {
		return "POC"
	}
}

func getStatus(s uint) string {
	if s == 0 {
		return "已过期"
	} else {
		return "维保中"
	}
}

func getLicenseStatus(l uint) string {
	if l == 1 {
		return "生效中"
	} else {
		return "已作废"
	}
}

func getMT(m uint) string {
	if m == 1 {
		return "付费"
	} else {
		return "免费"
	}
}

// GetCustomerSummary 获取客户汇总
func (cs *CustomerService) GetCustomerSummary() (customerData customerRes.CustomerData, err error) {
	var total int64
	// 获取客户总数
	customerDB := global.TD27_DB.Model(&CustomerModel.CustomerModel{}).Count(&total)
	customerData.Total.Customer = total

	// 获取客户类型数量
	var stats []customerRes.CustomerType
	customerDB.Select("type, COUNT(*) AS count").
		Group("type").
		Scan(&stats)
	for _, stat := range stats {
		item := customerRes.AnalyseItem{
			Value: stat.Count,
			Name:  getTypeName(stat.Type),
		}
		customerData.Analyse.CustomerType = append(customerData.Analyse.CustomerType, item)
	}

	// 获取维保是否过期的数量
	var ST []customerRes.TechnicalSupport
	global.TD27_DB.Model(&CustomerModel.CustomerModel{}).
		Select("service_status, COUNT(*) AS count").
		Group("service_status").
		Scan(&ST)

	for _, s := range ST {
		item := customerRes.AnalyseItem{
			Value: s.Count,
			Name:  getStatus(s.Status),
		}
		customerData.Analyse.TechnicalSupport = append(customerData.Analyse.TechnicalSupport, item)
	}

	// 获取付费/收费维保数量
	var MT []customerRes.MaintenanceType
	global.TD27_DB.Model(&CustomerModel.CustomerModel{}).
		Select("maintenance_type, COUNT(*) AS count").
		Group("maintenance_type").
		Scan(&MT)

	for _, s := range MT {
		if s.Type != 0 {
			item := customerRes.AnalyseItem{
				Value: s.Count,
				Name:  getMT(s.Type),
			}
			customerData.Analyse.MaintenanceType = append(customerData.Analyse.MaintenanceType, item)
		}
	}

	// 获取环境总数
	envDB := global.TD27_DB.Model(&CustomerModel.CustomerEnvModel{}).Count(&total)
	customerData.Total.Env = total
	// 获取不同类型的环境数量
	var env []customerRes.Env
	envDB.Select("env, COUNT(*) AS count").
		Group("env").
		Scan(&env)
	for _, e := range env {
		item := customerRes.AnalyseItem{
			Value: e.Count,
			Name:  e.Env,
		}
		customerData.Analyse.Env = append(customerData.Analyse.Env, item)
	}

	// 获取license总数
	licenseDB := global.TD27_DB.Model(&LicenseModel.LicenseModel{}).Count(&total)
	customerData.Total.License = total
	// 获取生效中和已作废的license数量
	var license []customerRes.License
	licenseDB.Select("active, COUNT(*) AS count").
		Group("active").
		Scan(&license)
	for _, l := range license {
		item := customerRes.AnalyseItem{
			Value: l.Count,
			Name:  getLicenseStatus(l.Active),
		}
		customerData.Analyse.License = append(customerData.Analyse.License, item)
	}

	return
}
