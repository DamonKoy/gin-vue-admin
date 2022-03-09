package device

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 返回数据库表名
type DeviceRecord struct {
	global.GVA_MODEL
	DeviceId       int64  `json:"device_id" gorm:"comment:设备id"`
	OriginalHolder string `json:"original_holder" gorm:"comment:原持有人"`
	CurrentHolder  string `json:"current_holder"  gorm:"comment:现持有人"`
	Describe       string `json:"describe" gorm:"comment:描述内容"`
}

func (DeviceRecord) TableName() string {
	return "device_transfer_record"
}
