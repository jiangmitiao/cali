package api

import "github.com/jiangmitiao/cali/app/services"

var (
	userService       = services.DefaultUserService
	userRoleService   = services.DefaultUserRoleService
	tagService        = services.DefaultTagService
	authorService     = services.DefaultAuthorService
	bookService       = services.DefaultBookService
	languageService   = services.DefaultLanguageService
	roleActionService = services.DefaultRoleActionService
	sysConfigService  = services.DefaultSysConfigService
)
