package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
)

type DeviceResponse struct {
	Device device.Device `json:"device"`
}
