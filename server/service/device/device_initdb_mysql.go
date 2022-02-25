package device

//
//import (
//	"github.com/flipped-aurora/gin-vue-admin/server/global"
//	"github.com/flipped-aurora/gin-vue-admin/server/model/device/response"
//)
//
//var DeviceMysql = new(deviceMysql)
//
//type deviceMysql struct{}
//
//// GetTables 获取手机设备列表
//func (s *deviceMysql) GetDeviceList() (data []response.DeviceResponse, err error) {
//	var entities []response.DeviceResponse
//	sql := `select * from device`
//	err = global.GVA_DB.Raw(sql).Scan(&entities).Error
//	return entities, err
//}
