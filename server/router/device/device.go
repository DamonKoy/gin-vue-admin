package device

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct{}

// 初始化路由信息
func (s *DeviceRouter) InitDeviceRouter(Router *gin.RouterGroup) {
	//deviceRouter := Router.Group("device").Use(middleware.OperationRecord())
	deviceRouterWithoutRecord := Router.Group("device")
	deviceApi := v1.ApiGroupApp.DeviceApiGroup.DeviceApi
	//{
	//	deviceRouter.POST("updateDeviceHolder", deviceApi.UpdateDeviceHolder) // 更新手机设备持有人
	//}
	{
		deviceRouterWithoutRecord.POST("getDeviceList", deviceApi.GetDeviceList) // 获取手机设备列表
	}
}
