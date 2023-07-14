package license

import (
	"server/service"
)

type ApiGroup struct {
	LicenseApi
}

var (
	licenseService = service.ServiceGroupApp.LicenseServiceGroup.LicenseService
)
