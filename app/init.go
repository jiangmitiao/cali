package app

import (
	"errors"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string = "v0.0.2"

	// BuildTime revel app build-time (ldflags)
	BuildTime string = time.Now().Format("2006-01-02")
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		rcali.QueryParamsFilter,
		revel.SessionFilter,    // Restore and write the session cookie.
		revel.FlashFilter,      // Restore and write the flash cookie.
		revel.ValidationFilter, // Restore kept validation errors and save new ones from cookie.
		//revel.I18nFilter,        // Resolve the requested language
		rcali.I18nFilter,
		HeaderFilter,            // Add some security based headers
		revel.InterceptorFilter, // Run interceptors around the action.
		revel.CompressFilter,    // Compress the result.
		revel.ActionInvoker,     // Invoke the action.
	}

	// register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)

	revel.OnAppStart(InitDebug)
	revel.OnAppStart(InitDB)
	//revel.OnAppStart(Monitor)
}

// HeaderFilter adds common security headers
// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}

func Monitor() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
}

//init db on first view
func InitDB() {
	if mysqlEnable, mysqlEnableFund := rcali.MysqlEnable(); mysqlEnableFund && mysqlEnable {
		//mysql
		mysqlDsn, _ := rcali.MysqlDsn()
		if err := services.DbInitByMysql(mysqlDsn); err != nil {
			panic(err)
		}
		defer services.UpdateSqlite2Mysql()
	} else if dbPath, dbPathFund := rcali.GetSqliteDbPath(); dbPathFund {
		if err := services.DbInitBySqlite(dbPath); err != nil {
			panic(err)
		}
	} else {
		panic(errors.New("sqlite path not found"))
	}

	if ok, err := services.DbInit(); ok {
		rcali.Logger.Info("---------------------dbok---------------------")
	} else {
		panic(err)
	}
}

func InitDebug() {
	//init the debug on first view page
	rcali.Logger = rcali.Log(revel.RunMode)
}
