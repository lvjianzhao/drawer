package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type ApiRouter struct{}

func (u *ApiRouter) InitApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiRouter := Router.Group("api")
	apiApi := api.ApiGroupApp.SystemApiGroup.ApiApi
	{
		apiRouter.POST("addApi", apiApi.AddApi)
		apiRouter.GET("getApis", apiApi.GetApis)
		apiRouter.GET("apiGroups", apiApi.GetApiGroups)
		apiRouter.GET("apiMethods", apiApi.GetApiMethods)
		apiRouter.POST("deleteApi", apiApi.DeleteApi)
		apiRouter.PUT("editApi", apiApi.EditApi)
		apiRouter.POST("getElTreeApis", apiApi.GetElTreeApis)
	}
	return apiRouter
}
