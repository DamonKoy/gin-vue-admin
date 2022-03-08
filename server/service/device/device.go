package device

import (
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/model/device/request"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
)

type DeviceService struct{}

func (deviceService *DeviceService) GetDeviceList(info deviceReq.GetDeviceList) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&device.Device{})
	var deviceList []device.Device
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&deviceList).Error
	return err, deviceList, total
}

func (deviceService *DeviceService) UpdateDeviceHolder(info deviceReq.UpdateDeviceHolder) (err error) {
	db := global.GVA_DB.Model(&device.Device{})
	var device []device.Device
	err = db.Where("id = ? and holder = ?", info.Id, info.Holder).First(&device).Update("holder", info.CurrentHolder).Error
	return err
}
