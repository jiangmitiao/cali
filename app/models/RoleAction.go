package models

type RoleAction struct {
	Id         string `json:"id" xorm:"pk 'id'"`
	Role       string `json:"role" xorm:"varchar(32) 'role'"`
	Controller string `json:"controller" xorm:"'controller'"`
	Method     string `json:"method" xorm:"'method'"`
	Comments   string `json:"comments" xorm:"'comments'"` //zhu shi
}

func (RoleAction) TableName() string {
	return "role_action"
}

var RoleActions = make([]RoleAction, 0)

//初始化的角色权限,watcher不可以下载
func init() {
	//admin
	RoleActions = append(RoleActions,
		RoleAction{Id: "adminBookIndex", Role: "admin", Controller: "Book", Method: "Index"},
		RoleAction{Id: "adminBookBooksCount", Role: "admin", Controller: "Book", Method: "BooksCount"},
		RoleAction{Id: "adminBookBooks", Role: "admin", Controller: "Book", Method: "Books"},
		RoleAction{Id: "adminBookBookDown", Role: "admin", Controller: "Book", Method: "BookDown"},
		RoleAction{Id: "adminBookBook", Role: "admin", Controller: "Book", Method: "Book"},
		RoleAction{Id: "adminBookDoubanBook", Role: "admin", Controller: "Book", Method: "DoubanBook"},
		RoleAction{Id: "adminBookUploadBook", Role: "admin", Controller: "Book", Method: "UploadBook"},
		RoleAction{Id: "adminBookUploadBookConfirm", Role: "admin", Controller: "Book", Method: "UploadBookConfirm"},
		RoleAction{Id: "adminBookSearch", Role: "admin", Controller: "Book", Method: "Search"},
		RoleAction{Id: "adminBookSearchCount", Role: "admin", Controller: "Book", Method: "SearchCount"},
		RoleAction{Id: "adminBookDelJSON", Role: "admin", Controller: "Book", Method: "DelJSON"},

		RoleAction{Id: "adminCategoryIndex", Role: "admin", Controller: "Category", Method: "Index"},
		RoleAction{Id: "adminCategoryAll", Role: "admin", Controller: "Category", Method: "All"},
		RoleAction{Id: "adminCategoryAdd", Role: "admin", Controller: "Category", Method: "Add"},
		RoleAction{Id: "adminCategoryUpdate", Role: "admin", Controller: "Category", Method: "Update"},
		RoleAction{Id: "adminCategoryDelete", Role: "admin", Controller: "Category", Method: "Delete"},

		RoleAction{Id: "adminUserIndex", Role: "admin", Controller: "User", Method: "Index"},
		RoleAction{Id: "adminUserLogin", Role: "admin", Controller: "User", Method: "Login"},
		RoleAction{Id: "adminUserInfo", Role: "admin", Controller: "User", Method: "Info"},
		RoleAction{Id: "adminUserIsLogin", Role: "admin", Controller: "User", Method: "IsLogin"},
		RoleAction{Id: "adminUserLogout", Role: "admin", Controller: "User", Method: "Logout"},
		RoleAction{Id: "adminUserRegist", Role: "admin", Controller: "User", Method: "Regist"},
		RoleAction{Id: "adminUserUpdate", Role: "admin", Controller: "User", Method: "Update"},
		RoleAction{Id: "adminUserChangePassword", Role: "admin", Controller: "User", Method: "ChangePassword"},
		RoleAction{Id: "adminUserQueryUserCount", Role: "admin", Controller: "User", Method: "QueryUserCount"},
		RoleAction{Id: "adminUserQueryUser", Role: "admin", Controller: "User", Method: "QueryUser"},
		RoleAction{Id: "adminUserChangeDelete", Role: "admin", Controller: "User", Method: "Delete"},
		RoleAction{Id: "adminUserUserStatus", Role: "admin", Controller: "User", Method: "UserStatus"},

		RoleAction{Id: "adminAppIndex", Role: "admin", Controller: "App", Method: "Index"},

		RoleAction{Id: "adminSysConfigIndex", Role: "admin", Controller: "SysConfig", Method: "Index"},
		RoleAction{Id: "adminSysConfigConfigs", Role: "admin", Controller: "SysConfig", Method: "Configs"},
		RoleAction{Id: "adminSysConfigUpdate", Role: "admin", Controller: "SysConfig", Method: "Update"},

		RoleAction{Id: "adminSysStatusIndex", Role: "admin", Controller: "SysStatus", Method: "Index"},
		RoleAction{Id: "adminSysStatusStatus", Role: "admin", Controller: "SysStatus", Method: "Status"},
		RoleAction{Id: "adminSysStatusDelete", Role: "admin", Controller: "SysStatus", Method: "Delete"},
	)

	//user
	RoleActions = append(RoleActions,
		RoleAction{Id: "userBookIndex", Role: "user", Controller: "Book", Method: "Index"},
		RoleAction{Id: "userBookBooksCount", Role: "user", Controller: "Book", Method: "BooksCount"},
		RoleAction{Id: "userBookBooks", Role: "user", Controller: "Book", Method: "Books"},
		RoleAction{Id: "userBookBookDown", Role: "user", Controller: "Book", Method: "BookDown"},
		RoleAction{Id: "userBookBook", Role: "user", Controller: "Book", Method: "Book"},
		RoleAction{Id: "userBookDoubanBook", Role: "user", Controller: "Book", Method: "DoubanBook"},
		RoleAction{Id: "userBookUploadBook", Role: "user", Controller: "Book", Method: "UploadBook"},
		RoleAction{Id: "userBookUploadBookConfirm", Role: "user", Controller: "Book", Method: "UploadBookConfirm"},
		RoleAction{Id: "userBookSearch", Role: "user", Controller: "Book", Method: "Search"},
		RoleAction{Id: "userBookSearchCount", Role: "user", Controller: "Book", Method: "SearchCount"},

		RoleAction{Id: "userCategoryIndex", Role: "user", Controller: "Category", Method: "Index"},
		RoleAction{Id: "userCategoryAll", Role: "user", Controller: "Category", Method: "All"},

		RoleAction{Id: "userUserIndex", Role: "user", Controller: "User", Method: "Index"},
		RoleAction{Id: "userUserLogin", Role: "user", Controller: "User", Method: "Login"},
		RoleAction{Id: "userUserInfo", Role: "user", Controller: "User", Method: "Info"},
		RoleAction{Id: "userUserIsLogin", Role: "user", Controller: "User", Method: "IsLogin"},
		RoleAction{Id: "userUserLogout", Role: "user", Controller: "User", Method: "Logout"},
		RoleAction{Id: "userUserRegist", Role: "user", Controller: "User", Method: "Regist"},
		RoleAction{Id: "userUserUpdate", Role: "user", Controller: "User", Method: "Update"},
		RoleAction{Id: "userUserChangePassword", Role: "user", Controller: "User", Method: "ChangePassword"},
		RoleAction{Id: "userUserUserStatus", Role: "user", Controller: "User", Method: "UserStatus"},

		RoleAction{Id: "userAppIndex", Role: "user", Controller: "App", Method: "Index"},
	)

	//watcher
	RoleActions = append(RoleActions,
		RoleAction{Id: "watcherBookIndex", Role: "watcher", Controller: "Book", Method: "Index"},
		RoleAction{Id: "watcherBookBooksCount", Role: "watcher", Controller: "Book", Method: "BooksCount"},
		RoleAction{Id: "watcherBookBooks", Role: "watcher", Controller: "Book", Method: "Books"},
		//RoleAction{Id: "watcherBookIndex", Role: "watcher", Controller: "Book", Method: "BookDown"},
		RoleAction{Id: "watcherBookBook", Role: "watcher", Controller: "Book", Method: "Book"},
		RoleAction{Id: "watcherBookDoubanBook", Role: "watcher", Controller: "Book", Method: "DoubanBook"},
		//RoleAction{Id: "watcherBookIndex", Role: "watcher", Controller: "Book", Method: "UploadBook"},
		RoleAction{Id: "watcherBookSearch", Role: "watcher", Controller: "Book", Method: "Search"},
		RoleAction{Id: "watcherBookSearchCount", Role: "watcher", Controller: "Book", Method: "SearchCount"},

		RoleAction{Id: "watcherCategoryIndex", Role: "watcher", Controller: "Category", Method: "Index"},
		RoleAction{Id: "watcherCategoryAll", Role: "watcher", Controller: "Category", Method: "All"},

		RoleAction{Id: "watcherUserIndex", Role: "watcher", Controller: "User", Method: "Index"},
		RoleAction{Id: "watcherUserLogin", Role: "watcher", Controller: "User", Method: "Login"},
		RoleAction{Id: "watcherUserInfo", Role: "watcher", Controller: "User", Method: "Info"},
		RoleAction{Id: "watcherUserIsLogin", Role: "watcher", Controller: "User", Method: "IsLogin"},
		RoleAction{Id: "watcherUserLogout", Role: "watcher", Controller: "User", Method: "Logout"},
		RoleAction{Id: "watcherUserRegist", Role: "watcher", Controller: "User", Method: "Regist"},
		//RoleAction{Id: "watcherUserUpdate", Role: "watcher", Controller: "User", Method: "Update"}, watcher can not change info
		//RoleAction{Id: "watcherUserChangePassword", Role: "watcher", Controller: "User", Method: "ChangePassword"}, watcher can not change password
		RoleAction{Id: "watcherUserActive", Role: "watcher", Controller: "User", Method: "Active"},

		RoleAction{Id: "watcherAppIndex", Role: "watcher", Controller: "App", Method: "Index"},
	)
}
