package initialize

import (
	"source_manager_admin_api/global"
	"source_manager_admin_api/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(&model.Houses{},&model.Images{})
	global.GVA_LOG.Debug("register table success")
}
