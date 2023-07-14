package license

import (
	"encoding/json"
	"fmt"
	"server/global"
	CustomerModel "server/model/customer"
	LicenseModel "server/model/license"
	licenseReq "server/model/license/request"
	licenseRes "server/model/license/response"
	"server/tools"
)

type LicenseService struct{}

type GenLicenseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		LicStr string `json:"licStr"`
	} `json:"data"`
}

// AddLicense 添加license
func (ls *LicenseService) AddLicense(license licenseReq.License) (*licenseRes.LicenseResult, error) {
	var licenseModel LicenseModel.LicenseModel
	var licenseResult licenseRes.LicenseResult

	// 创建新的 LicenseModel 对象
	licenseModel = LicenseModel.LicenseModel{
		CustomerID:           license.CustomerID,
		Env:                  license.ENV,
		ExpiredDate:          tools.ParseDateTime(license.ExpiredDate),
		ManagerApis:          license.ManagerApis,
		ManagerUsers:         license.ManagerUsers,
		ManagerProjects:      license.ManagerProjects,
		ManagerIntegrations:  license.ManagerIntegrations,
		ManagerOrganizations: license.ManagerOrganizations,
		ManagerTests:         license.ManagerTests,
		ManagerProducts:      license.ManagerProducts,
		StudioUsers:          license.StudioUsers,
		StudioCompBlackList:  license.StudioCompBlackList,
		PermissionMenu:       license.PermissionMenu,
		CheckList:            license.CheckList,
	}

	licenseCfg := global.TD27_CONFIG.License
	// 根据customerID获取客户名称
	var CompName string
	err := global.TD27_DB.Model(&CustomerModel.CustomerModel{}).
		Select("name").
		Where("id = ?", license.CustomerID).
		First(&CompName).
		Error
	if err != nil {
		return nil, err
	}

	CompFullName := CompName + license.ENV

	license.CompName = tools.ZhongwenToPinyin(CompName)
	license.CompFullName = tools.ZhongwenToPinyin(CompFullName)

	requestBody, err := tools.StructToJsonStr(license)
	if err != nil {
		return nil, err
	}

	res, err := tools.MakeAPIRequest(licenseCfg.Auth, licenseCfg.Url, "POST", requestBody, licenseCfg.Authorization)
	if err != nil {
		return nil, err
	}

	var licenseData GenLicenseRes
	err = json.Unmarshal(res, &licenseData)
	if err != nil {
		global.TD27_LOG.Info(fmt.Sprintf("创建license接口返回的数据: %s", res))
		global.TD27_LOG.Error(fmt.Sprintf("解析创建license接口返回的数据出错,error: %s", err))
	}

	// 更新licenseModel的License字段
	licenseModel.License = licenseData.Data.LicStr
	// 把之前的license置为失效
	db := global.TD27_DB.Model(&LicenseModel.LicenseModel{}).
		Where("customer_id = ? AND env = ?", license.CustomerID, license.ENV).
		Update("active", 2)

	if db.Error != nil {
		global.TD27_LOG.Error(fmt.Sprintf("license置为失效失败,error: %s", err))
	}

	// 执行添加操作
	if err := global.TD27_DB.Create(&licenseModel).Error; err != nil {
		return nil, err
	}

	//// 构造返回结果
	licenseResult.ID = licenseModel.ID
	licenseResult.CreatedAt = licenseModel.CreatedAt
	licenseResult.License = licenseModel.License
	return &licenseResult, nil
}

// GetLicense 获取license列表
func (ls *LicenseService) GetLicense(licenseReq licenseReq.LicenseSearchParams) ([]licenseRes.LicenseResult, int64, error) {
	var licenseResults []licenseRes.LicenseResult
	var total int64

	db := global.TD27_DB.Model(&LicenseModel.LicenseModel{}).
		Select("license.*, customer.name").
		Joins("JOIN customer ON license.customer_id = customer.id").
		Order("license.active ASC").
		Order("license.created_at DESC")

	if licenseReq.CustomerName != "" {
		db = db.Where("customer.name LIKE ?", "%"+licenseReq.CustomerName+"%")
	}

	if licenseReq.Active != 0 {
		db = db.Where("license.active = ?", licenseReq.Active)
	}
	
	err := db.Count(&total).Error
	db.Limit(licenseReq.PageSize).Offset(licenseReq.PageSize * (licenseReq.Page - 1)).Find(&licenseResults)

	if err != nil {
		return licenseResults, total, fmt.Errorf("查询 LicenseModel 失败: %v", err)
	}
	// 返回结果
	return licenseResults, total, nil
}

// DeleteLicense 删除客户license记录
func (ls *LicenseService) DeleteLicense(id int) (err error) {
	return global.TD27_DB.Where("id = ?", id).Unscoped().Delete(&LicenseModel.LicenseModel{}).Error
}
