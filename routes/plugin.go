package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"panel/app/http/controllers/plugins/supervisor"

	"panel/app/http/controllers/plugins/mysql57"
	"panel/app/http/controllers/plugins/mysql80"
	"panel/app/http/controllers/plugins/openresty"
	"panel/app/http/controllers/plugins/php74"
	"panel/app/http/controllers/plugins/php80"
	"panel/app/http/controllers/plugins/phpmyadmin"
	"panel/app/http/controllers/plugins/s3fs"
	"panel/app/http/middleware"
)

// Plugin 加载插件路由
func Plugin() {
	facades.Route().Prefix("api/plugins/openresty").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		openRestyController := openresty.NewOpenrestyController()
		route.Get("status", openRestyController.Status)
		route.Post("reload", openRestyController.Reload)
		route.Post("start", openRestyController.Start)
		route.Post("stop", openRestyController.Stop)
		route.Post("restart", openRestyController.Restart)
		route.Get("load", openRestyController.Load)
		route.Get("config", openRestyController.GetConfig)
		route.Post("config", openRestyController.SaveConfig)
		route.Get("errorLog", openRestyController.ErrorLog)
		route.Post("clearErrorLog", openRestyController.ClearErrorLog)
	})
	facades.Route().Prefix("api/plugins/mysql57").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		mysql57Controller := mysql57.NewMysql57Controller()
		route.Get("status", mysql57Controller.Status)
		route.Post("reload", mysql57Controller.Reload)
		route.Post("start", mysql57Controller.Start)
		route.Post("stop", mysql57Controller.Stop)
		route.Post("restart", mysql57Controller.Restart)
		route.Get("load", mysql57Controller.Load)
		route.Get("config", mysql57Controller.GetConfig)
		route.Post("config", mysql57Controller.SaveConfig)
		route.Get("errorLog", mysql57Controller.ErrorLog)
		route.Post("clearErrorLog", mysql57Controller.ClearErrorLog)
		route.Get("slowLog", mysql57Controller.SlowLog)
		route.Post("clearSlowLog", mysql57Controller.ClearSlowLog)
		route.Get("rootPassword", mysql57Controller.GetRootPassword)
		route.Post("rootPassword", mysql57Controller.SetRootPassword)
		route.Get("database", mysql57Controller.DatabaseList)
		route.Post("addDatabase", mysql57Controller.AddDatabase)
		route.Post("deleteDatabase", mysql57Controller.DeleteDatabase)
		route.Get("backup", mysql57Controller.BackupList)
		route.Post("createBackup", mysql57Controller.CreateBackup)
		route.Post("uploadBackup", mysql57Controller.UploadBackup)
		route.Post("deleteBackup", mysql57Controller.DeleteBackup)
		route.Post("restoreBackup", mysql57Controller.RestoreBackup)
		route.Get("user", mysql57Controller.UserList)
		route.Post("addUser", mysql57Controller.AddUser)
		route.Post("deleteUser", mysql57Controller.DeleteUser)
		route.Post("userPassword", mysql57Controller.SetUserPassword)
		route.Post("userPrivileges", mysql57Controller.SetUserPrivileges)
	})
	facades.Route().Prefix("api/plugins/mysql80").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		mysql80Controller := mysql80.NewMysql80Controller()
		route.Get("status", mysql80Controller.Status)
		route.Post("reload", mysql80Controller.Reload)
		route.Post("start", mysql80Controller.Start)
		route.Post("stop", mysql80Controller.Stop)
		route.Post("restart", mysql80Controller.Restart)
		route.Get("load", mysql80Controller.Load)
		route.Get("config", mysql80Controller.GetConfig)
		route.Post("config", mysql80Controller.SaveConfig)
		route.Get("errorLog", mysql80Controller.ErrorLog)
		route.Post("clearErrorLog", mysql80Controller.ClearErrorLog)
		route.Get("slowLog", mysql80Controller.SlowLog)
		route.Post("clearSlowLog", mysql80Controller.ClearSlowLog)
		route.Get("rootPassword", mysql80Controller.GetRootPassword)
		route.Post("rootPassword", mysql80Controller.SetRootPassword)
		route.Get("database", mysql80Controller.DatabaseList)
		route.Post("addDatabase", mysql80Controller.AddDatabase)
		route.Post("deleteDatabase", mysql80Controller.DeleteDatabase)
		route.Get("backup", mysql80Controller.BackupList)
		route.Post("createBackup", mysql80Controller.CreateBackup)
		route.Post("uploadBackup", mysql80Controller.UploadBackup)
		route.Post("deleteBackup", mysql80Controller.DeleteBackup)
		route.Post("restoreBackup", mysql80Controller.RestoreBackup)
		route.Get("user", mysql80Controller.UserList)
		route.Post("addUser", mysql80Controller.AddUser)
		route.Post("deleteUser", mysql80Controller.DeleteUser)
		route.Post("userPassword", mysql80Controller.SetUserPassword)
		route.Post("userPrivileges", mysql80Controller.SetUserPrivileges)
	})
	facades.Route().Prefix("api/plugins/php74").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		php74Controller := php74.NewPhp74Controller()
		route.Get("status", php74Controller.Status)
		route.Post("reload", php74Controller.Reload)
		route.Post("start", php74Controller.Start)
		route.Post("stop", php74Controller.Stop)
		route.Post("restart", php74Controller.Restart)
		route.Get("load", php74Controller.Load)
		route.Get("config", php74Controller.GetConfig)
		route.Post("config", php74Controller.SaveConfig)
		route.Get("errorLog", php74Controller.ErrorLog)
		route.Get("slowLog", php74Controller.SlowLog)
		route.Post("clearErrorLog", php74Controller.ClearErrorLog)
		route.Post("clearSlowLog", php74Controller.ClearSlowLog)
		route.Get("extensions", php74Controller.GetExtensionList)
		route.Post("installExtension", php74Controller.InstallExtension)
		route.Post("uninstallExtension", php74Controller.UninstallExtension)
	})
	facades.Route().Prefix("api/plugins/php80").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		php80Controller := php80.NewPhp80Controller()
		route.Get("status", php80Controller.Status)
		route.Post("reload", php80Controller.Reload)
		route.Post("start", php80Controller.Start)
		route.Post("stop", php80Controller.Stop)
		route.Post("restart", php80Controller.Restart)
		route.Get("load", php80Controller.Load)
		route.Get("config", php80Controller.GetConfig)
		route.Post("config", php80Controller.SaveConfig)
		route.Get("errorLog", php80Controller.ErrorLog)
		route.Get("slowLog", php80Controller.SlowLog)
		route.Post("clearErrorLog", php80Controller.ClearErrorLog)
		route.Post("clearSlowLog", php80Controller.ClearSlowLog)
		route.Get("extensions", php80Controller.GetExtensionList)
		route.Post("installExtension", php80Controller.InstallExtension)
		route.Post("uninstallExtension", php80Controller.UninstallExtension)
	})
	facades.Route().Prefix("api/plugins/phpmyadmin").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		phpMyAdminController := phpmyadmin.NewPhpMyAdminController()
		route.Get("info", phpMyAdminController.Info)
		route.Post("port", phpMyAdminController.SetPort)
	})
	facades.Route().Prefix("api/plugins/s3fs").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		s3fsController := s3fs.NewS3fsController()
		route.Get("list", s3fsController.List)
		route.Post("add", s3fsController.Add)
		route.Post("delete", s3fsController.Delete)
	})
	facades.Route().Prefix("api/plugins/supervisor").Middleware(middleware.Jwt()).Group(func(route route.Route) {
		supervisorController := supervisor.NewSupervisorController()
		route.Get("status", supervisorController.Status)
		route.Post("start", supervisorController.Start)
		route.Post("stop", supervisorController.Stop)
		route.Post("restart", supervisorController.Restart)
		route.Post("reload", supervisorController.Reload)
		route.Get("log", supervisorController.Log)
		route.Post("clearLog", supervisorController.ClearLog)
		route.Get("config", supervisorController.Config)
		route.Post("config", supervisorController.SaveConfig)
		route.Get("processes", supervisorController.Processes)
		route.Post("startProcess", supervisorController.StartProcess)
		route.Post("stopProcess", supervisorController.StopProcess)
		route.Post("restartProcess", supervisorController.RestartProcess)
		route.Get("processLog", supervisorController.ProcessLog)
		route.Post("clearProcessLog", supervisorController.ClearProcessLog)
		route.Get("processConfig", supervisorController.ProcessConfig)
		route.Post("processConfig", supervisorController.SaveProcessConfig)
		route.Post("deleteProcess", supervisorController.DeleteProcess)
		route.Post("addProcess", supervisorController.AddProcess)

	})
}
