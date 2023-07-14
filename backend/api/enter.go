package api

import (
	"server/api/customer"
	"server/api/license"
	"server/api/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	CustomerGroup  customer.ApiGroup
	LicenseGroup   license.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
