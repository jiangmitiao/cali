package api

import "github.com/jiangmitiao/cali/app/services"

var (
	bookService   = services.DefaultBookService
	formatService = services.DefaultFormatService
	categoryService = services.DefaultCategoryService

	userService       = services.DefaultUserService
	userRoleService   = services.DefaultUserRoleService
	roleActionService = services.DefaultRoleActionService
	sysConfigService  = services.DefaultSysConfigService
	sysStatusService  = services.DefaultSysStatusService
	userConfigService = services.DefaultUserConfigService
)
