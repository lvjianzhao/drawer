package service

import (
	"server/service/customer"
	"server/service/license"
	"server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	CustomerServiceGroup customer.ServiceGroup
	LicenseServiceGroup  license.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
