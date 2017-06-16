package interceptor

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
	"strings"
)

var (
	userService       = services.DefaultUserService
	userRoleService   = services.DefaultUserRoleService
	roleActionService = services.DefaultRoleActionService
	sysConfigService  = services.DefaultSysConfigService

	//roleActionCache controller action role
	roleActionCache = make(map[string]map[string]map[string]string)
)

func validateOK(controller, method, role string) bool {

	if methods, ok := roleActionCache[controller]; ok {
		if roles, ok := methods[method]; ok {
			if allow, ok := roles[role]; ok {
				if allow != "" {
					return true
				} else {
					rcali.Logger.Debug("this action need to login :", controller, method, role)
					return false
				}
			}
		}
	}

	roleAction := roleActionService.GetRoleActionByControllerMethodRole(controller, method, role)
	if roleActionCache[controller] == nil {
		roleActionCache[controller] = make(map[string]map[string]string)
	}
	if roleActionCache[controller][method] == nil {
		roleActionCache[controller][method] = make(map[string]string)
	}
	roleActionCache[controller][method][role] = roleAction.Id
	if roleAction.Id == "" {
		rcali.Logger.Debug("this action need to login :", controller, method, role)
		return false
	} else {
		return true
	}
}

func authInterceptor(c *revel.Controller) revel.Result {
	// 全部变成首字大写
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if controller == "Static" { //不拦截静态地址
		return nil
	}
	if controller == "View" {
		return nil
	}
	rcali.Logger.Debug("controller: ", controller)
	rcali.Logger.Debug("method: ", method)

	session := c.Request.Form.Get("session")

	id, _ := rcali.GetUserIdByLoginSession(session)
	var role models.Role
	if id != "" {
		role = userRoleService.GetRoleByUser(id)
	}
	roleId := role.Id
	if roleId == "" {
		roleId = "watcher"
	}

	rcali.Logger.Debug("role: ", roleId)

	if validateOK(controller, method, roleId) {
		return nil
	} else {
		return c.Redirect("/login")
	}
}

func configInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	if controller == "View" {
		c.ViewArgs["cnzzid"] = sysConfigService.Get("cnzzid")
		return nil
	}
	return nil
}

func init() {
	revel.InterceptFunc(authInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(configInterceptor, revel.AFTER, revel.AllControllers)
}
