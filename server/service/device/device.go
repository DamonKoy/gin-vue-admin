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
	// 更新device表的holder值
	deviceDb := global.GVA_DB.Model(&device.Device{})
	var deviceTable []device.Device
	err = deviceDb.Where("id = ? and holder = ?", info.Id, info.Holder).First(&deviceTable).Update("Holder", info.CurrentHolder).Error
	if err != nil {
		return err
	} else {
		// 在device_transfer_record表新增一条记录
		deviceRecordDb := global.GVA_DB.Model(&device.DeviceRecord{})
		record := deviceReq.DeviceRecord{DeviceId: info.Id, OriginalHolder: info.Holder, CurrentHolder: info.CurrentHolder, Describe: info.Describe}
		err = deviceRecordDb.Create(&record).Error
		return err
	}
}
