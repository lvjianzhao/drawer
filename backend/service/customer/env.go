package customer

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	CustomerModel "server/model/customer"
	customerReq "server/model/customer/request"
	customerRes "server/model/customer/response"
)

// GetCustomerEnv 获取客户环境列表
func (cs *CustomerService) GetCustomerEnv(EnvSp customerReq.CustomerSearchParams) ([]customerRes.CustomerEnvResult, int64, error) {
	var envResults []customerRes.CustomerEnvResult
	var total int64

	if EnvSp.ID != 0 {
		err := global.TD27_DB.Model(&CustomerModel.CustomerEnvModel{}).
			Select("env,customer_id").
			Where("customer_id = ?", EnvSp.ID).
			Order("updated_at DESC").
			Find(&envResults).Error
		return envResults, total, err
	}

	db := global.TD27_DB.Model(&CustomerModel.CustomerEnvModel{}).
		Select("customer_env.*, customer.name").
		Joins("JOIN customer ON customer_env.customer_id = customer.id").
		Order("customer_env.updated_at DESC")

	if EnvSp.Name != "" {
		db = db.Where("customer.name LIKE ?", "%"+EnvSp.Name+"%")
	}

	if EnvSp.IP != "" {
		db = db.Where("customer_env.deploy_report LIKE ?", "%"+EnvSp.IP+"%")
	}
	err := db.Count(&total).Error
	db.Limit(EnvSp.PageSize).Offset(EnvSp.PageSize * (EnvSp.Page - 1)).Find(&envResults)
	if err != nil {
		return envResults, total, fmt.Errorf("查询 CustomerEnvModel 失败: %v", err)
	}
	// 返回结果
	return envResults, total, nil
}

// AddOrUpdateCustomerEnv 添加/更新客户环境信息
func (cs *CustomerService) AddOrUpdateCustomerEnv(env customerReq.CustomerEnv) (*customerRes.CustomerEnvResult, string, error) {
	var customerEnvModel CustomerModel.CustomerEnvModel
	var customerEnvResult customerRes.CustomerEnvResult
	var action string

	if env.ID == 0 {
		action = "添加"
		// 创建新的 CustomerEnvModel 对象
		customerEnvModel = CustomerModel.CustomerEnvModel{
			CustomerID:   env.CustomerId,
			Env:          env.Env,
			Vpn:          env.Vpn,
			DeployReport: env.DeployReport,
		}

		// 执行添加操作
		if err := global.TD27_DB.Create(&customerEnvModel).Error; err != nil {
			return nil, action, err
		}
	} else {
		action = "修改"
		//// 根据 env.ID 查询对应的 customerEnvModel 对象
		if err := global.TD27_DB.First(&customerEnvModel, env.ID).Error; err != nil {
			global.TD27_LOG.Error("编辑客户环境信息 -> 查询Id", zap.Error(err))
			return nil, action, err
		}
		//
		//// 更新 customerEnvModel 对象的属性值
		customerEnvModel.CustomerID = env.CustomerId
		customerEnvModel.Env = env.Env
		customerEnvModel.Vpn = env.Vpn
		customerEnvModel.DeployReport = env.DeployReport

		// 执行修改操作
		if err := global.TD27_DB.Save(&customerEnvModel).Error; err != nil {
			global.TD27_LOG.Error("编辑客户环境 -> update", zap.Error(err))
			return nil, action, err
		}
	}

	// 构造返回结果
	customerEnvResult.ID = customerEnvModel.ID
	customerEnvResult.UpdatedAt = customerEnvModel.UpdatedAt
	customerEnvResult.CustomerId = customerEnvModel.CustomerID
	customerEnvResult.Env = customerEnvModel.Env

	return &customerEnvResult, action, nil
}

// DeleteCustomerEnv 删除客户环境
func (cs *CustomerService) DeleteCustomerEnv(id int) (err error) {
	return global.TD27_DB.Where("id = ?", id).Unscoped().Delete(&CustomerModel.CustomerEnvModel{}).Error
}
