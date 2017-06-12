package interceptor

import (
	"github.com/jiangmitiao/cali/app/rcali"
	"github.com/jiangmitiao/cali/app/services"
	"github.com/revel/revel"
	"strings"
)

var (
	dbok = false
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

// 拦截器
// true is need login
// false is not need login
var commonUrl = map[string]map[string]bool{
	"Book": map[string]bool{
		"Index":              false,
		"BooksCount":         false,
		"Books":              false,
		"RatingBooks":        false,
		"NewBooks":           false,
		"DiscoverBooks":      false,
		"TagBooksCount":      false,
		"TagBooks":           false,
		"AuthorBooksCount":   false,
		"AuthorBooks":        false,
		"LanguageBooksCount": false,
		"LanguageBooks":      false,
		"BookRating":         false,
		"BookImage":          false,
		"BookDown":           true, // download need login
		"Book":               false,
		"DoubanBook":         false,
		"UploadBook":         true, // upload need login
	},
	"Author": map[string]bool{
		"Index":        false,
		"AuthorsCount": false,
		"Authors":      false,
	},
	"Language": map[string]bool{
		"Index":          false,
		"LanguagesCount": false,
		"Languages":      false,
	},
	"Tag": map[string]bool{
		"Index":     false,
		"TagsCount": false,
		"Tags":      false,
	},
	"User": map[string]bool{
		"Index":  false,
		"Login":  false,
		"Logout": true, //logout need loged
		"Regist": false,
	},
}

func needValidate(controller, method string) bool {
	if controller == "Static" { //不拦截静态地址
		return false
	}
	// 在里面
	if v, ok := commonUrl[controller]; ok {
		// 在commonUrl里
		if need, ok2 := v[method]; ok2 {
			return need
		}
		return true
	} else {
		// controller不在这里的, 肯定要验证
		return true
	}
}

func authInterceeptor(c *revel.Controller) revel.Result {
	// 全部变成首字大写
	var controller = strings.Title(c.Name)
	var method = strings.Title(c.MethodName)
	rcali.DEBUG.Debug("controller", controller)
	rcali.DEBUG.Debug("method", method)

	if needValidate(controller, method) {
		var session string
		c.Params.Bind(&session, "session")
		id, _ := rcali.GetUserIdByLoginSession(session)
		rcali.DEBUG.Debug("session: " + session + " id: " + id)
		return c.Redirect("/public/v/login.html")
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
	revel.InterceptFunc(authInterceeptor, revel.BEFORE, revel.AllControllers)
}
