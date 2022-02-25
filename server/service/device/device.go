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

//func (deviceService *DeviceService) GetDeviceList(info request.PageInfo) (err error, deviceList interface{}, total int64) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	db := global.GVA_DB.Model(&device.Device{})
//	var deviceList []device.Device
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//	err = db.Limit(limit).Offset(offset).Preload("device").Preload("device").Find(&deviceList).Error
//	return err, deviceList, total
//}

//func (deviceService *DeviceService) updateDeviceHolder(info request.holder) (err error, list interface{}, total int64) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	db := global.GVA_DB.Model(&device.Device{})
//	var deviceList []device.Device
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//	err = db.Limit(limit).Offset(offset).Preload("Device").Preload("Device").Find(&deviceList).Error
//	return err, deviceList, total
//}
