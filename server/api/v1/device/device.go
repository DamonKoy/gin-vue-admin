package device

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"
)

type DeviceApi struct{}

//func (deviceApi *DeviceApi) updateDeviceHolder(c *gin.Context) {
//	var device device.Device
//	_ = c.ShouldBindJSON(&device)
//	if err := deviceService.updateDeviceHolder(&device); err != nil {
//		global.GVA_LOG.Error("更新失败!", zap.Error(err))
//		response.FailWithMessage("更新失败", c)
//	} else {
//		response.OkWithMessage("更新成功", c)
//	}
//}

func (b *DeviceApi) GetDeviceList(c *gin.Context) {
	var pageInfo deviceReq.GetDeviceList
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, deviceList, total := deviceService.GetDeviceList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     deviceList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
