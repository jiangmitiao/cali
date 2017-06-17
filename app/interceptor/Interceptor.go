package interceptor

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/juju/ratelimit"
	"github.com/revel/revel"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	userService       = services.DefaultUserService
	userRoleService   = services.DefaultUserRoleService
	roleActionService = services.DefaultRoleActionService
	sysConfigService  = services.DefaultSysConfigService

	//roleActionCache controller action role
	roleActionCache = make(map[string]map[string]map[string]string)

	//download need to limit
	limitLock         = sync.Mutex{}
	limitTokenBuckets = make(map[string]*ratelimit.Bucket)
)

func roleActionGet(controller, method, role string) string {
	//roleActionCache
	if methods, ok := roleActionCache[controller]; ok {
		if roles, ok := methods[method]; ok {
			if roleid, ok := roles[role]; ok {
				return roleid
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
	return roleActionCache[controller][method][role]

}

func validateOK(controller, method, role string) bool {
	roleid := roleActionGet(controller, method, role)
	if roleid == "" {
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

	id, _ := rcali.GetUserIdByLoginSession(c.Request.Form.Get("session"))
	var role models.Role
	if id != "" {
		role = userRoleService.GetRoleByUser(id)
	}
	roleId := role.Id
	if roleId == "" {
		roleId = "watcher"
	}

	c.Params.Set("roleid", roleId)

	rcali.Logger.Debug("role: ", roleId)

	if validateOK(controller, method, roleId) {
		return nil
	} else {
		return c.Redirect("/login")
	}
}

func openRegistInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if (controller == "View" && method == "SignUp") || (controller == "User" && method == "Regist") {
		//allow register?
		if sysConfigService.Get("openregist") == "true" {
			return nil
		} else {
			return c.Redirect("/")
		}
	}
	return nil
}

func configInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	if controller == "View" {
		c.ViewArgs["cnzzid"] = sysConfigService.Get("cnzzid")
		return nil
	}
	return nil
}

func takeAvailable(userId string, maxDayLimit int64) int64 {
	limitLock.Lock()
	tokenBucket, ok := limitTokenBuckets[userId]
	limitLock.Unlock()

	if !ok {
		tokenBucket = ratelimit.NewBucket(time.Hour*24, maxDayLimit)
		limitLock.Lock()
		limitTokenBuckets[userId] = tokenBucket
		limitLock.Unlock()
	}
	return limitTokenBuckets[userId].TakeAvailable(1)
}

//download action need to limit ,to defense attack http://blog.imlibo.com/2016/06/20/golang-token-bucket/
func downloadLimitInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if controller == "Book" && method == "BookDown" {
		limitConfig,_ := strconv.Atoi(sysConfigService.Get("alldownloadlimit"))
		if takeAvailable("common", int64(limitConfig)) <= 0 {
			return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApiWithMessageAndInfo(c.Message("limitdownload"), nil))
		}
	}
	return nil
}

func init() {
	revel.InterceptFunc(authInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(openRegistInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(downloadLimitInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(configInterceptor, revel.AFTER, revel.AllControllers)

}
