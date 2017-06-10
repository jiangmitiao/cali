package interceptor

import (
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
)

var (
	dbok = false
)

//init db on first view
func dbInterceptor(c *revel.Controller) revel.Result {
	if c.Name == "Static" { //不拦截静态地址
		return nil
	}

	if dbok && c.Action == "Install.Install" {
		c.Redirect("/")
	}
	if c.Action == "Install.Index" {
		return nil
	}

	if !dbok {
		//加载db
		if sqlitedbpath, ok := rcali.GetSqliteDbPath(); ok {
			rcali.DEBUG.Debug("database init " + sqlitedbpath)
			if ok, _ := services.DbInit(sqlitedbpath); ok {
				dbok = true
				return nil
			} else {
				return c.Redirect("install/")
			}
		} else {
			return c.Redirect("install/")
		}

	}

	return nil
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
}
