package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
)

// 返回数据库表名
type DeviceRecordResponse struct {
	DeviceRecord device.DeviceRecord `json:"device_record"`
}
