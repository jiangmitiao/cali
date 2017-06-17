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
		RoleAction{Id: "adminBookRatingBooks", Role: "admin", Controller: "Book", Method: "RatingBooks"},
		RoleAction{Id: "adminBookNewBooks", Role: "admin", Controller: "Book", Method: "NewBooks"},
		RoleAction{Id: "adminBookDiscoverBooks", Role: "admin", Controller: "Book", Method: "DiscoverBooks"},
		RoleAction{Id: "adminBookTagBooksCount", Role: "admin", Controller: "Book", Method: "TagBooksCount"},
		RoleAction{Id: "adminBookTagBooks", Role: "admin", Controller: "Book", Method: "TagBooks"},
		RoleAction{Id: "adminBookAuthorBooksCount", Role: "admin", Controller: "Book", Method: "AuthorBooksCount"},
		RoleAction{Id: "adminBookAuthorBooks", Role: "admin", Controller: "Book", Method: "AuthorBooks"},
		RoleAction{Id: "adminBookLanguageBooksCount", Role: "admin", Controller: "Book", Method: "LanguageBooksCount"},
		RoleAction{Id: "adminBookLanguageBooks", Role: "admin", Controller: "Book", Method: "LanguageBooks"},
		RoleAction{Id: "adminBookBookRating", Role: "admin", Controller: "Book", Method: "BookRating"},
		RoleAction{Id: "adminBookBookImage", Role: "admin", Controller: "Book", Method: "BookImage"},
		RoleAction{Id: "adminBookBookDown", Role: "admin", Controller: "Book", Method: "BookDown"},
		RoleAction{Id: "adminBookBook", Role: "admin", Controller: "Book", Method: "Book"},
		RoleAction{Id: "adminBookDoubanBook", Role: "admin", Controller: "Book", Method: "DoubanBook"},
		RoleAction{Id: "adminBookUploadBook", Role: "admin", Controller: "Book", Method: "UploadBook"},
		RoleAction{Id: "adminBookSearch", Role: "admin", Controller: "Book", Method: "Search"},
		RoleAction{Id: "adminBookSearchCount", Role: "admin", Controller: "Book", Method: "SearchCount"},

		RoleAction{Id: "adminAuthorIndex", Role: "admin", Controller: "Author", Method: "Index"},
		RoleAction{Id: "adminAuthorAuthorsCount", Role: "admin", Controller: "Author", Method: "AuthorsCount"},
		RoleAction{Id: "adminAuthorAuthors", Role: "admin", Controller: "Author", Method: "Authors"},

		RoleAction{Id: "adminLanguageIndex", Role: "admin", Controller: "Language", Method: "Index"},
		RoleAction{Id: "adminLanguageLanguagesCount", Role: "admin", Controller: "Language", Method: "LanguagesCount"},
		RoleAction{Id: "adminLanguageLanguages", Role: "admin", Controller: "Language", Method: "Languages"},

		RoleAction{Id: "adminTagIndex", Role: "admin", Controller: "Tag", Method: "Index"},
		RoleAction{Id: "adminTagTagsCount", Role: "admin", Controller: "Tag", Method: "TagsCount"},
		RoleAction{Id: "adminTagTags", Role: "admin", Controller: "Tag", Method: "Tags"},

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

		RoleAction{Id: "adminAppIndex", Role: "admin", Controller: "App", Method: "Index"},

		RoleAction{Id: "adminSysConfigIndex", Role: "admin", Controller: "SysConfig", Method: "Index"},
		RoleAction{Id: "adminSysConfigConfigs", Role: "admin", Controller: "SysConfig", Method: "Configs"},
		RoleAction{Id: "adminSysConfigUpdate", Role: "admin", Controller: "SysConfig", Method: "Update"},
	)

	//user
	RoleActions = append(RoleActions,
		RoleAction{Id: "userBookIndex", Role: "user", Controller: "Book", Method: "Index"},
		RoleAction{Id: "userBookBooksCount", Role: "user", Controller: "Book", Method: "BooksCount"},
		RoleAction{Id: "userBookBooks", Role: "user", Controller: "Book", Method: "Books"},
		RoleAction{Id: "userBookRatingBooks", Role: "user", Controller: "Book", Method: "RatingBooks"},
		RoleAction{Id: "userBookNewBooks", Role: "user", Controller: "Book", Method: "NewBooks"},
		RoleAction{Id: "userBookDiscoverBooks", Role: "user", Controller: "Book", Method: "DiscoverBooks"},
		RoleAction{Id: "userBookTagBooksCount", Role: "user", Controller: "Book", Method: "TagBooksCount"},
		RoleAction{Id: "userBookTagBooks", Role: "user", Controller: "Book", Method: "TagBooks"},
		RoleAction{Id: "userBookAuthorBooksCount", Role: "user", Controller: "Book", Method: "AuthorBooksCount"},
		RoleAction{Id: "userBookAuthorBooks", Role: "user", Controller: "Book", Method: "AuthorBooks"},
		RoleAction{Id: "userBookLanguageBooksCount", Role: "user", Controller: "Book", Method: "LanguageBooksCount"},
		RoleAction{Id: "userBookLanguageBooks", Role: "user", Controller: "Book", Method: "LanguageBooks"},
		RoleAction{Id: "userBookBookRating", Role: "user", Controller: "Book", Method: "BookRating"},
		RoleAction{Id: "userBookBookImage", Role: "user", Controller: "Book", Method: "BookImage"},
		RoleAction{Id: "userBookBookDown", Role: "user", Controller: "Book", Method: "BookDown"},
		RoleAction{Id: "userBookBook", Role: "user", Controller: "Book", Method: "Book"},
		RoleAction{Id: "userBookDoubanBook", Role: "user", Controller: "Book", Method: "DoubanBook"},
		RoleAction{Id: "userBookUploadBook", Role: "user", Controller: "Book", Method: "UploadBook"},
		RoleAction{Id: "userBookSearch", Role: "user", Controller: "Book", Method: "Search"},
		RoleAction{Id: "userBookSearchCount", Role: "user", Controller: "Book", Method: "SearchCount"},

		RoleAction{Id: "userAuthorIndex", Role: "user", Controller: "Author", Method: "Index"},
		RoleAction{Id: "userAuthorAuthorsCount", Role: "user", Controller: "Author", Method: "AuthorsCount"},
		RoleAction{Id: "userAuthorAuthors", Role: "user", Controller: "Author", Method: "Authors"},

		RoleAction{Id: "userLanguageIndex", Role: "user", Controller: "Language", Method: "Index"},
		RoleAction{Id: "userLanguageLanguagesCount", Role: "user", Controller: "Language", Method: "LanguagesCount"},
		RoleAction{Id: "userLanguageLanguages", Role: "user", Controller: "Language", Method: "Languages"},

		RoleAction{Id: "userTagIndex", Role: "user", Controller: "Tag", Method: "Index"},
		RoleAction{Id: "userTagTagsCount", Role: "user", Controller: "Tag", Method: "TagsCount"},
		RoleAction{Id: "userTagTags", Role: "user", Controller: "Tag", Method: "Tags"},

		RoleAction{Id: "userUserIndex", Role: "user", Controller: "User", Method: "Index"},
		RoleAction{Id: "userUserLogin", Role: "user", Controller: "User", Method: "Login"},
		RoleAction{Id: "userUserInfo", Role: "user", Controller: "User", Method: "Info"},
		RoleAction{Id: "userUserIsLogin", Role: "user", Controller: "User", Method: "IsLogin"},
		RoleAction{Id: "userUserLogout", Role: "user", Controller: "User", Method: "Logout"},
		RoleAction{Id: "userUserRegist", Role: "user", Controller: "User", Method: "Regist"},
		RoleAction{Id: "userUserUpdate", Role: "user", Controller: "User", Method: "Update"},
		RoleAction{Id: "userUserChangePassword", Role: "user", Controller: "User", Method: "ChangePassword"},

		RoleAction{Id: "userAppIndex", Role: "user", Controller: "App", Method: "Index"},
	)

	//watcher
	RoleActions = append(RoleActions,
		RoleAction{Id: "watcherBookIndex", Role: "watcher", Controller: "Book", Method: "Index"},
		RoleAction{Id: "watcherBookBooksCount", Role: "watcher", Controller: "Book", Method: "BooksCount"},
		RoleAction{Id: "watcherBookBooks", Role: "watcher", Controller: "Book", Method: "Books"},
		RoleAction{Id: "watcherBookRatingBooks", Role: "watcher", Controller: "Book", Method: "RatingBooks"},
		RoleAction{Id: "watcherBookNewBooks", Role: "watcher", Controller: "Book", Method: "NewBooks"},
		RoleAction{Id: "watcherBookDiscoverBooks", Role: "watcher", Controller: "Book", Method: "DiscoverBooks"},
		RoleAction{Id: "watcherBookTagBooksCount", Role: "watcher", Controller: "Book", Method: "TagBooksCount"},
		RoleAction{Id: "watcherBookTagBooks", Role: "watcher", Controller: "Book", Method: "TagBooks"},
		RoleAction{Id: "watcherBookAuthorBooksCount", Role: "watcher", Controller: "Book", Method: "AuthorBooksCount"},
		RoleAction{Id: "watcherBookAuthorBooks", Role: "watcher", Controller: "Book", Method: "AuthorBooks"},
		RoleAction{Id: "watcherBookLanguageBooksCount", Role: "watcher", Controller: "Book", Method: "LanguageBooksCount"},
		RoleAction{Id: "watcherBookLanguageBooks", Role: "watcher", Controller: "Book", Method: "LanguageBooks"},
		RoleAction{Id: "watcherBookBookRating", Role: "watcher", Controller: "Book", Method: "BookRating"},
		RoleAction{Id: "watcherBookBookImage", Role: "watcher", Controller: "Book", Method: "BookImage"},
		//RoleAction{Id: "watcherBookIndex", Role: "watcher", Controller: "Book", Method: "BookDown"},
		RoleAction{Id: "watcherBookBook", Role: "watcher", Controller: "Book", Method: "Book"},
		RoleAction{Id: "watcherBookDoubanBook", Role: "watcher", Controller: "Book", Method: "DoubanBook"},
		//RoleAction{Id: "watcherBookIndex", Role: "watcher", Controller: "Book", Method: "UploadBook"},
		RoleAction{Id: "watcherBookSearch", Role: "watcher", Controller: "Book", Method: "Search"},
		RoleAction{Id: "watcherBookSearchCount", Role: "watcher", Controller: "Book", Method: "SearchCount"},


		RoleAction{Id: "watcherAuthorIndex", Role: "watcher", Controller: "Author", Method: "Index"},
		RoleAction{Id: "watcherAuthorAuthorsCount", Role: "watcher", Controller: "Author", Method: "AuthorsCount"},
		RoleAction{Id: "watcherAuthorAuthors", Role: "watcher", Controller: "Author", Method: "Authors"},

		RoleAction{Id: "watcherLanguageIndex", Role: "watcher", Controller: "Language", Method: "Index"},
		RoleAction{Id: "watcherLanguageLanguagesCount", Role: "watcher", Controller: "Language", Method: "LanguagesCount"},
		RoleAction{Id: "watcherLanguageLanguages", Role: "watcher", Controller: "Language", Method: "Languages"},

		RoleAction{Id: "watcherTagIndex", Role: "watcher", Controller: "Tag", Method: "Index"},
		RoleAction{Id: "watcherTagTagsCount", Role: "watcher", Controller: "Tag", Method: "TagsCount"},
		RoleAction{Id: "watcherTagTags", Role: "watcher", Controller: "Tag", Method: "Tags"},

		RoleAction{Id: "watcherUserIndex", Role: "watcher", Controller: "User", Method: "Index"},
		RoleAction{Id: "watcherUserLogin", Role: "watcher", Controller: "User", Method: "Login"},
		RoleAction{Id: "watcherUserInfo", Role: "watcher", Controller: "User", Method: "Info"},
		RoleAction{Id: "watcherUserIsLogin", Role: "watcher", Controller: "User", Method: "IsLogin"},
		RoleAction{Id: "watcherUserLogout", Role: "watcher", Controller: "User", Method: "Logout"},
		RoleAction{Id: "watcherUserRegist", Role: "watcher", Controller: "User", Method: "Regist"},
		//RoleAction{Id: "watcherUserUpdate", Role: "watcher", Controller: "User", Method: "Update"}, watcher can not change info
		//RoleAction{Id: "watcherUserChangePassword", Role: "watcher", Controller: "User", Method: "ChangePassword"}, watcher can not change password

		RoleAction{Id: "watcherAppIndex", Role: "watcher", Controller: "App", Method: "Index"},
	)
}
