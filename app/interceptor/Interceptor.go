package interceptor

import (
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/juju/ratelimit"
	"github.com/revel/revel"
	"github.com/shirou/gopsutil/mem"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"runtime"
)

var (
	userService       = services.DefaultUserService
	userRoleService   = services.DefaultUserRoleService
	roleActionService = services.DefaultRoleActionService
	sysConfigService  = services.DefaultSysConfigService
	sysStatusService  = services.DefaultSysStatusService
	userConfigService = services.DefaultUserConfigService

	//roleActionCache controller action role
	roleActionCache     = make(map[string]map[string]map[string]string)
	roleActionCacheLock = sync.Mutex{}

	//download need to limit
	limitLock         = sync.Mutex{}
	limitTokenBuckets = make(map[string]*ratelimit.Bucket)
)

func roleActionGet(controller, method, role string) string {
	roleActionCacheLock.Lock()
	defer roleActionCacheLock.Unlock()
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
		if sysConfigService.Get("openregist").Value == "true" {
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
		c.ViewArgs["cnzzid"] = sysConfigService.Get("cnzzid").Value
		return nil
	}
	return nil
}

func takeAvailable(key string, maxLimit int64, gap time.Duration) int64 {
	if maxLimit <= 1 {
		maxLimit = 1
	}
	limitLock.Lock()
	tokenBucket, ok := limitTokenBuckets[key]
	limitLock.Unlock()

	realGap := time.Duration(int64(gap.Nanoseconds() / maxLimit))

	if !ok {
		tokenBucket = ratelimit.NewBucket(realGap, maxLimit)
		limitLock.Lock()
		limitTokenBuckets[key] = tokenBucket
		limitLock.Unlock()
	} else {
		//changed then
		if tokenBucket.Capacity() != maxLimit {
			newTokenBucket := ratelimit.NewBucket(realGap, maxLimit)
			limitLock.Lock()
			limitTokenBuckets[key] = newTokenBucket
			limitLock.Unlock()
			newTokenBucket.TakeAvailable(tokenBucket.Available())
		}
	}
	//not allow to download
	if maxLimit == 1 {
		limitTokenBuckets[key].TakeAvailable(1)
	}
	return limitTokenBuckets[key].TakeAvailable(1)
}

//download action need to limit ,to defense attack http://blog.imlibo.com/2016/06/20/golang-token-bucket/
func downloadLimitInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if controller == "Book" && method == "BookDown" {
		limitConfig, _ := strconv.Atoi(sysConfigService.Get("alldownloadlimit").Value)
		if takeAvailable("common", int64(limitConfig), time.Hour*24) <= 0 {
			return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApiWithMessageAndInfo(c.Message("limitdownload"), nil))
		}
		// add status to sys status
		key := time.Now().Format("20060102") + "-downnum"
		if status := sysStatusService.Get(key); status.Key != "" {
			value, _ := strconv.ParseInt(status.Value, 10, 0)
			value += 1
			status.Value = strconv.FormatInt(value, 10)
			sysStatusService.UpdateStatus(status)
		} else {
			status = models.SysStatus{Key: key, Value: strconv.FormatInt(1, 10)}
			sysStatusService.AddSysStatus(status)
		}
	}

	return nil
}

func userDownloadLimitInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if controller == "Book" && method == "BookDown" {
		tmp := time.Now()
		start := time.Date(tmp.Year(), tmp.Month(), tmp.Day(), 0, 0, 0, 0, tmp.Location())
		stop := start.AddDate(0, 0, 1)
		user, _ := userService.GetLoginUser(c.Request.FormValue("session"))
		if user.Id == "" {
			return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApiWithMessageAndInfo(c.Message("limitdownload"), nil))
		}
		count := userService.GetDownloadCount(user.Id, start, stop)
		config, _ := userConfigService.GetUserConfig(user.Id)
		if count >= config.MaxDownload {
			return c.RenderJSONP(c.Request.FormValue("callback"), models.NewErrorApiWithMessageAndInfo(c.Message("limitdownload"), nil))
		}
	}

	return nil
}

func ipLimitInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	limitConfig, _ := strconv.Atoi(sysConfigService.Get("iplimit").Value)
	if method == "BookImage" {
		limitConfig = limitConfig * 8
	}
	for takeAvailable(c.ClientIP+controller+method, int64(limitConfig), time.Minute) <= 0 {
		time.Sleep(time.Second * 1)
	}
	return nil
}

func sysStatusInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	//var method = strings.Title(c.MethodName)
	if controller == "View" {
		// add status to sys status
		key := time.Now().Format("20060102") + "-pv"
		if status := sysStatusService.Get(key); status.Value != "" {
			pvi, _ := strconv.ParseInt(status.Value, 10, 0)
			status.Value = strconv.FormatInt(pvi+1, 10)
			sysStatusService.UpdateStatus(status)
		} else {
			sysStatusService.AddSysStatus(models.SysStatus{Key: key, Value: strconv.Itoa(1)})
		}
	}
	return nil
}

func monitorInterceptor(c *revel.Controller) revel.Result {
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	if controller == "View" && method == "Person" {
		v, _ := mem.VirtualMemory()
		c, _ := cpu.Info()
		cc, _ := cpu.Percent(time.Second, false)
		d, _ := disk.Usage("/")
		n, _ := host.Info()
		nv, _ := net.IOCounters(true)
		boottime, _ := host.BootTime()
		btime := time.Unix(int64(boottime), 0).Format("2006-01-02 15:04:05")

		fmt.Printf("        Mem       : %v MB  Free: %v MB Used:%v Usage:%f%%\n", v.Total/1024/1024, v.Available/1024/1024, v.Used/1024/1024, v.UsedPercent)
		if len(c) > 1 {
			for _, sub_cpu := range c {
				modelname := sub_cpu.ModelName
				cores := sub_cpu.Cores
				fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)
			}
		} else {
			sub_cpu := c[0]
			modelname := sub_cpu.ModelName
			cores := sub_cpu.Cores
			fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)

		}
		fmt.Printf("        Network: %v bytes / %v bytes\n", nv[0].BytesRecv, nv[0].BytesSent)
		fmt.Printf("        SystemBoot:%v\n", btime)
		fmt.Printf("        CPU Used    : used %f%% \n", cc[0])
		fmt.Printf("        HD        : %v GB  Free: %v GB Usage:%f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
		fmt.Printf("        OS        : %v(%v)   %v  \n", n.Platform, n.PlatformFamily, n.PlatformVersion)
		fmt.Printf("        Hostname  : %v  \n", n.Hostname)
	}

	runtime.GC()
	return nil
}

func init() {
	//revel.InterceptFunc(ipLimitInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(authInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(userDownloadLimitInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(openRegistInterceptor, revel.BEFORE, revel.AllControllers)
	revel.InterceptFunc(downloadLimitInterceptor, revel.AFTER, revel.AllControllers)
	revel.InterceptFunc(configInterceptor, revel.AFTER, revel.AllControllers)
	revel.InterceptFunc(sysStatusInterceptor, revel.AFTER, revel.AllControllers)

	revel.InterceptFunc(monitorInterceptor, revel.AFTER, revel.AllControllers)

}
