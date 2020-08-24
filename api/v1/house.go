package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"source_manager_admin_api/global/response"
	"source_manager_admin_api/model/request"
	resp "source_manager_admin_api/model/response"
	"source_manager_admin_api/service"
	"source_manager_admin_api/utils"
)

// @Tags HouseApi
// @Summary 获取二手房列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "获取二手房列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/list [get]
func GetHouseFang(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}

	err, list, total := service.GetHousesInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
