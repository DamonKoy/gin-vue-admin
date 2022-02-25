package device

/*
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"go.uber.org/zap"
)

var DeviceMysql = new(deviceMysql)

type deviceMysql struct{}

// GetColumn 获取指定数据库和指定数据表的所有字段名和值
func (s *deviceMysql) GetColumn(tableName string, dbName string) (data []response.Column, err error) {
	global.GVA_LOG.Debug("我在device_mysql")
	var entities []response.Column
	sql := `SELECT id,model,platform,system_version,auditor,holder,update_time,remarks FROM device `
	global.GVA_LOG.Error(sql, zap.Error(err))
	err = global.GVA_DB.Raw(sql, tableName, dbName).Scan(&entities).Error
	return entities, err
}
*/
