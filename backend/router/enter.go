package router

import (
	"server/router/customer"
	"server/router/license"
	"server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Customer customer.RouterGroup
	License  license.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
