package customer

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	CustomerModel "server/model/customer"
	customerReq "server/model/customer/request"
	customerRes "server/model/customer/response"
	LicenseModel "server/model/license"
	"server/tools"
)

type CustomerService struct{}

type allCustomer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetAllCustomer 获取客户列表
func (cs *CustomerService) GetAllCustomer() ([]allCustomer, int64, error) {
	var allCustomerResults []allCustomer
	var total int64

	db := global.TD27_DB.Model(&CustomerModel.CustomerModel{})

	err := db.Count(&total).Error
	db.Pluck("id,name", &allCustomerResults)

	return allCustomerResults, total, err
}

// GetCustomer 获取客户列表
func (cs *CustomerService) GetCustomer(CustomerSp customerReq.CustomerSearchParams) ([]customerRes.CustomerResult, int64, error) {
	var customerResults []customerRes.CustomerResult
	var total int64

	db := global.TD27_DB.Model(&CustomerModel.CustomerModel{})

	if CustomerSp.Name != "" {
		db = db.Where("name LIKE ?", "%"+CustomerSp.Name+"%")
	}

	// 0: 已过期, 1:维保中,2: 所有状态
	if CustomerSp.ServiceStatus != 2 {
		db = db.Where("service_status = ?", CustomerSp.ServiceStatus)
	}

	if CustomerSp.Type != 0 {
		db = db.Where("type = ?", CustomerSp.Type)
	}

	// go里面0和空值没有区别，所以需要二次判断并重新赋值, 前端如果传的值是3，则表示要查询maintenance_type为0的记录
	if CustomerSp.MaintenanceType != 0 {
		if CustomerSp.MaintenanceType == 3 {
			CustomerSp.MaintenanceType = 0
		}
		db = db.Where("maintenance_type = ?", CustomerSp.MaintenanceType)
	}

	// 分页
	err := db.Count(&total).Error
	if err != nil {
		return customerResults, total, fmt.Errorf("分页count -> %v", err)
	} else {
		limit := CustomerSp.PageSize
		offset := CustomerSp.PageSize * (CustomerSp.Page - 1)

		db = db.Limit(limit).Offset(offset)

		// 指定排序条件并查询客户信息
		db = db.Order("updated_at DESC").Find(&customerResults)
	}

	return customerResults, total, err
}

// AddOrUpdateCustomer 添加/更新客户
func (cs *CustomerService) AddOrUpdateCustomer(customer customerReq.Customer) (*customerRes.CustomerResult, string, error) {
	var customerModel CustomerModel.CustomerModel
	var customerResult customerRes.CustomerResult
	var action string

	if customer.ID == 0 {
		action = "添加"
		// 创建新的 CustomerModel 对象
		customerModel = CustomerModel.CustomerModel{
			Name:             customer.Name,
			Type:             customer.Type,
			BusinessName:     customer.BusinessName,
			Description:      customer.Description,
			ServiceStartDate: tools.ParseDate(customer.ServiceStartDate),
			ServiceEndDate:   tools.ParseDate(customer.ServiceEndDate),
			MaintenanceType:  customer.MaintenanceType,
		}

		// 执行添加操作
		if err := global.TD27_DB.Create(&customerModel).Error; err != nil {
			return nil, action, err
		}
	} else {
		action = "修改"
		// 根据 customer.ID 查询对应的 CustomerModel 对象
		if err := global.TD27_DB.First(&customerModel, customer.ID).Error; err != nil {
			global.TD27_LOG.Error("编辑客户名称 -> 查询Id", zap.Error(err))
			return nil, action, err
		}
		fmt.Println(customer)

		// 更新 CustomerModel 对象的属性值
		customerModel.Name = customer.Name
		customerModel.Type = customer.Type
		customerModel.BusinessName = customer.BusinessName
		customerModel.Description = customer.Description
		customerModel.ServiceStartDate = tools.ParseDate(customer.ServiceStartDate)
		customerModel.ServiceEndDate = tools.ParseDate(customer.ServiceEndDate)
		customerModel.MaintenanceType = customer.MaintenanceType

		// 执行修改操作
		if err := global.TD27_DB.Save(&customerModel).Error; err != nil {
			global.TD27_LOG.Error("编辑用户 -> update", zap.Error(err))
			return nil, action, err
		}
	}

	// 构造返回结果
	customerResult.ID = customerModel.ID
	customerResult.UpdatedAt = customerModel.UpdatedAt
	customerResult.Name = customerModel.Name
	customerResult.Type = customerModel.Type
	customerResult.BusinessName = customerModel.BusinessName
	customerResult.ServiceStartDate = customerModel.ServiceStartDate
	customerResult.ServiceEndDate = customerModel.ServiceEndDate
	customerResult.ServiceStatus = customerModel.ServiceStatus
	customerResult.Description = customerModel.Description
	customerResult.MaintenanceType = customerModel.MaintenanceType

	return &customerResult, action, nil
}

// DeleteCustomer 删除客户
func (cs *CustomerService) DeleteCustomer(id int) (err error) {
	tx := global.TD27_DB.Begin() // 开启事务

	// 删除 CustomerEnvModel 中关联的记录
	if err := tx.Where("customer_id = ?", id).Unscoped().Delete(&CustomerModel.CustomerEnvModel{}).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	// 删除 CustomerEventModel 中关联的记录
	if err := tx.Where("customer_id = ?", id).Unscoped().Delete(&CustomerModel.CustomerEventModel{}).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	// 删除 LicenseModel 中关联的记录
	if err := tx.Where("customer_id = ?", id).Unscoped().Delete(&LicenseModel.LicenseModel{}).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	// 删除 CustomerModel
	if err := tx.Where("id = ?", id).Unscoped().Delete(&CustomerModel.CustomerModel{}).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	tx.Commit() // 提交事务
	return nil
}
