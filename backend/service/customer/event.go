package customer

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	CustomerModel "server/model/customer"
	customerReq "server/model/customer/request"
	customerRes "server/model/customer/response"
	"server/tools"
	"time"
)

// AddOrUpdateEvent 添加/更新客户事件
func (cs *CustomerService) AddOrUpdateEvent(event customerReq.CustomerEvent) (*customerRes.CustomerEventResult, string, error) {
	var customerEventModel CustomerModel.CustomerEventModel
	var customerEventResult customerRes.CustomerEventResult
	var action string

	if event.ID == 0 {
		action = "添加"
		// 创建新的 customerEventModel 对象
		customerEventModel = CustomerModel.CustomerEventModel{
			CustomerID:   event.CustomerId,
			EventTime:    tools.ParseDateTime(event.EventTime),
			EventContent: event.EventContent,
		}

		// 执行添加操作
		if err := global.TD27_DB.Create(&customerEventModel).Error; err != nil {
			return nil, action, err
		}
	} else {
		action = "修改"
		//// 根据 event.ID 查询对应的 customerEventModel 对象
		if err := global.TD27_DB.First(&customerEventModel, event.ID).Error; err != nil {
			global.TD27_LOG.Error("编辑客户事件 -> 查询Id", zap.Error(err))
			return nil, action, err
		}
		//
		//// 更新 customerEventModel 对象的属性值
		customerEventModel.CustomerID = event.CustomerId
		customerEventModel.EventContent = event.EventContent
		customerEventModel.EventTime = tools.ParseDateTime(event.EventTime)

		// 执行修改操作
		if err := global.TD27_DB.Save(&customerEventModel).Error; err != nil {
			global.TD27_LOG.Error("编辑客户事件 -> update", zap.Error(err))
			return nil, action, err
		}
	}

	// 构造返回结果
	customerEventResult.ID = customerEventModel.ID
	customerEventResult.UpdatedAt = customerEventModel.UpdatedAt
	customerEventResult.CustomerId = customerEventModel.CustomerID
	customerEventResult.EventTime = customerEventModel.EventTime
	customerEventResult.EventContent = customerEventModel.EventContent

	return &customerEventResult, action, nil
}

// GetCustomerEvent 获取客户事件
func (cs *CustomerService) GetCustomerEvent(EventSp customerReq.EventSearchParams) ([]customerRes.CustomerEventResult, int64, error) {
	var eventResults []customerRes.CustomerEventResult
	var total int64

	db := global.TD27_DB.Model(&CustomerModel.CustomerEventModel{}).
		Select("customer_event.*, customer.name").
		Joins("JOIN customer ON customer_event.customer_id = customer.id")

	if EventSp.Sort == "eventTime" {
		db = db.Order("customer_event.event_time DESC")
	} else {
		db = db.Order("customer_event.updated_at DESC")
	}

	if EventSp.CustomerName != "" {
		db = db.Where("customer.name LIKE ?", "%"+EventSp.CustomerName+"%")
	}

	if EventSp.EventKeyword != "" {
		db = db.Where("customer_event.event_content LIKE ?", "%"+EventSp.EventKeyword+"%")
	}

	if EventSp.CustomerID != 0 {
		db = db.Where("customer_event.customer_id = ?", EventSp.CustomerID)
	}

	if EventSp.StartTime != 0 && EventSp.EndTime != 0 {
		startTime := time.Unix(EventSp.StartTime, 0)
		endTime := time.Unix(EventSp.EndTime, 0)
		db.Where("event_time BETWEEN ? AND ?", startTime, endTime)
	}

	err := db.Count(&total).Error
	db.Limit(EventSp.PageSize).Offset(EventSp.PageSize * (EventSp.Page - 1)).Find(&eventResults)
	if err != nil {
		return eventResults, total, fmt.Errorf("查询 CustomerEventModel 失败: %v", err)
	}
	// 返回结果
	return eventResults, total, nil
}

// DeleteEvent  删除客户事件
func (cs *CustomerService) DeleteEvent(id int) (err error) {
	return global.TD27_DB.Where("id = ?", id).Unscoped().Delete(&CustomerModel.CustomerEventModel{}).Error
}
