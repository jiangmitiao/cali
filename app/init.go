package app

import (
	"errors"
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
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
		revel.SessionFilter,     // Restore and write the session cookie.
		revel.FlashFilter,       // Restore and write the flash cookie.
		revel.ValidationFilter,  // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,        // Resolve the requested language
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

//init db on first view
func InitDB() {
	if dbPath, found := rcali.GetSqliteDbPath(); found {
		if ok, err := services.DbInit(dbPath); ok {
			rcali.Logger.Info("---------------------dbok---------------------")
		} else {
			panic(err)
		}
	} else {
		panic(errors.New("sqlite path not found"))
	}
}

func InitDebug() {
	//init the debug on first view page
	rcali.Logger = rcali.Log(revel.RunMode)
}
