package device

//
//import (
//	"database/sql"
//	"fmt"
//
//	"github.com/flipped-aurora/gin-vue-admin/server/global"
//	"github.com/flipped-aurora/gin-vue-admin/server/model/device"
//	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
//)
//
//type InitDBService struct{}
//
////// InitDB 创建数据库并初始化 总入口
////func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
////	switch conf.DBType {
////	case "mysql":
////		return initDBService.initMsqlDB(conf)
////	default:
////		return initDBService.initMsqlDB(conf)
////	}
////}
////
////// initTables 初始化表
////func (initDBService *InitDBService) initTables() error {
////	return global.GVA_DB.AutoMigrate(
////		device.Device{},
////	)
////}
////
////// createDatabase 创建数据库(mysql)
////func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
////	db, err := sql.Open(driver, dsn)
////	if err != nil {
////		return err
////	}
////	defer func(db *sql.DB) {
////		err = db.Close()
////		if err != nil {
////			fmt.Println(err)
////		}
////	}(db)
////	if err = db.Ping(); err != nil {
////		return err
////	}
////	_, err = db.Exec(createSql)
////	return err
////}
