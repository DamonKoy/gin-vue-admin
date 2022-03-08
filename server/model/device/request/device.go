package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

// device list structure
type GetDeviceList struct {
	request.PageInfo
}

type UpdateDeviceHolder struct {
	Id            int64  `json:"id" form:"id"`                       // id
	Holder        string `json:"holder" form:"holder"`               // 借出人
	CurrentHolder string `json:"currentHolder" form:"currentHolder"` // 借入人
	Remarks       string `json:"remarks" form:"remarks"`             // 备注信息
}
