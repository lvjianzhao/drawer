package customer

import (
	"server/service"
)

type ApiGroup struct {
	CustomerApi
}

var (
	customerService = service.ServiceGroupApp.CustomerServiceGroup.CustomerService
)
