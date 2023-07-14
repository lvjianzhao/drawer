package license

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type LicenseRouter struct{}

func (l *LicenseRouter) InitLicenseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	licenseRouter := Router.Group("license")
	licenseApi := api.ApiGroupApp.LicenseGroup.LicenseApi
	{
		// 客户维保信息
		licenseRouter.GET("", licenseApi.GetLicense)
		licenseRouter.POST("", licenseApi.AddLicense)
		licenseRouter.DELETE("/:licenseId", licenseApi.DeleteLicense)
	}
	return licenseRouter
}
