package device

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Device struct {
	global.GVA_MODEL
	Id            uint      `json:"id" gorm:"comment:设备id"`
	Icon          string    `json:"icon" gorm:"comment:icon"`
	Brand         string    `json:"brand" gorm:"comment:品牌"`
	Model         string    `json:"model"  gorm:"comment:型号"`
	Platform      string    `json:"platform" gorm:"comment:平台"`
	SystemVersion string    `json:"system_version" gorm:"comment:系统版本"`
	Resolution    string    `json:"resolution gorm:"`
	Auditor       string    `json:"auditor" gorm:"comment:管理人"`
	Holder        string    `json:"holder" gorm:"default:#fff;comment:持有人"`
	Status        string    `json:"status" gorm:"comment:状态"`
	Remarks       string    `json:"remarks" gorm:"comment:备注"`
	CreateTime    time.Time `json:"create_time" gorm:"comment:创建时间"`
	UpdateTime    time.Time `json:"update_time" gorm:"comment:更新时间"`
}

// 返回数据库表名
func (Device) TableName() string {
	return "device"
}
