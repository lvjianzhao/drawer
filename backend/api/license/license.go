package license

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	commReq "server/model/common/request"
	"server/model/common/response"
	licenseReq "server/model/license/request"
	"server/tools"
)

type LicenseApi struct{}

// AddLicense 添加License
func (la *LicenseApi) AddLicense(c *gin.Context) {
	var license licenseReq.License
	_ = c.ShouldBindJSON(&license)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&license); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if res, err := licenseService.AddLicense(license); err != nil {
		msg := fmt.Sprintf("失败, %s", err)
		response.FailWithMessage(msg, c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		msg := fmt.Sprintf("创建成功")
		response.OkWithDetailed(res, msg, c)
	}
}

// GetLicense 查询license
func (la *LicenseApi) GetLicense(c *gin.Context) {
	var licenseSp licenseReq.LicenseSearchParams
	licenseSp.Page, _ = tools.StringToInt(c.Query("currentPage"))
	licenseSp.PageSize, _ = tools.StringToInt(c.Query("size"))
	licenseSp.CustomerName = c.Query("name")
	licenseSp.Active, _ = tools.StringToUint(c.Query("active"))

	if list, total, err := licenseService.GetLicense(licenseSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取客户列表失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			Page:     licenseSp.Page,
			PageSize: licenseSp.PageSize,
			Total:    total,
			List:     list,
		}, "获取成功", c)
	}
}

func (la *LicenseApi) DeleteLicense(c *gin.Context) {
	id, _ := tools.StringToInt(c.Param("licenseId"))
	var cId commReq.CId
	_ = c.ShouldBindJSON(&cId)

	if err := licenseService.DeleteLicense(id); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
