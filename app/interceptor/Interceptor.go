package interceptor

import (
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
	"strings"
)

var (
	dbok = false

	userService       = services.DefaultUserService
	userRoleService   = services.DefaultUserRoleService
	roleActionService = services.DefaultRoleActionService
)

//init db on first view
func dbInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if controller == "Static" { //不拦截静态地址
		return nil
	}

	if dbok && controller == "Install" {
		return c.Redirect("/")
	}
	if controller == "Install" && method == "Index" {
		return nil
	}

	if !dbok {
		//加载db
		if sqlitedbpath, ok := rcali.GetSqliteDbPath(); ok {
			rcali.DEBUG.Debug("database init " + sqlitedbpath)
			if ok, err := services.DbInit(sqlitedbpath); ok {
				dbok = true
				return nil
			} else {
				rcali.DEBUG.Debug("database error ", err)
				return c.Redirect("install/")
			}
		} else {
			return c.Redirect("install/")
		}

	}
	return nil
}

func validateOK(controller, method, role string) bool {
	roleAction := roleActionService.GetRoleActionByControllerMethodRole(controller, method, role)
	if roleAction.Id == "" {
		rcali.DEBUG.Debug("this action need to login :", controller, method, role)
		return false
	} else {
		return true
	}
}

func authInterceeptor(c *revel.Controller) revel.Result {
	// 全部变成首字大写
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if controller == "Static" { //不拦截静态地址
		return nil
	}
	if controller == "Install" {
		return nil
	}
	var session string
	rcali.DEBUG.Debug("controller: ", controller)
	rcali.DEBUG.Debug("method: ", method)

	c.Params.Bind(&session, "session")
	if session == "" {
		session = c.Request.Form.Get("session")
	}
	id, _ := rcali.GetUserIdByLoginSession(session)
	role := userRoleService.GetRoleByUser(id)
	roleId := role.Id
	if roleId == "" {
		roleId = "watcher"
	}

	rcali.DEBUG.Debug("role: ", roleId)

	if validateOK(controller, method, roleId) {
		return nil
	} else {
		return c.Redirect("/public/v/login.html")
	}
}

//init the debug on first view page
func DebugInterceptor(c *revel.Controller) revel.Result {
	if rcali.DEBUG == "" {
		rcali.DEBUG = rcali.Debug(revel.RunMode)
		return nil
	} else {
		return nil
	}
}

func init() {
	revel.InterceptFunc(DebugInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(dbInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(authInterceeptor, revel.BEFORE, revel.AllControllers)
}
