package main

import (
	"source_manager_admin_api/core"
	"source_manager_admin_api/global"
	"source_manager_admin_api/initialize"
	//"runtime"
)

// @title Source-manage
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.Mysql()
	initialize.DBTables()
	// 程序结束前关闭数据库链接
	defer global.GVA_DB.Close()

	core.RunWindowsServer()
}
