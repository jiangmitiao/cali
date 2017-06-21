package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	//"github.com/google/uuid"
	"github.com/jiangmitiao/cali/app/models"
	"github.com/jiangmitiao/cali/app/rcali"
	_ "github.com/mattn/go-sqlite3"
	"path"
	"sync"
)

var (
	engine      *xorm.Engine
	localEngine *xorm.Engine

	DefaultAuthorService     = AuthorService{}
	DefaultBookService       = BookService{lock: &sync.Mutex{}}
	DefaultLanguageService   = LanguageService{}
	DefaultTagService        = TagService{}
	DefaultUserService       = UserService{}
	DefaultUserRoleService   = UserRoleService{}
	DefaultRoleActionService = RoleActionService{}
	DefaultSysConfigService  = SysConfigService{}
	DefaultSysStatusService  = SysStatusService{lock: &sync.Mutex{}}
	DefaultUserConfigService = UserConfigService{}
)

//init the db,should take a db filepath
func DbInit(SqliteDbPath string) (bool, error) { //username, password, host, database string
	if exist, err := rcali.FileExists(SqliteDbPath); !exist {
		rcali.Logger.Error("sqlitedbpath is error", SqliteDbPath, err)
		return false, err
	}

	var err error
	if engine, err = xorm.NewEngine("sqlite3", SqliteDbPath); err != nil {
		rcali.Logger.Error("open sqlitedb fail on ", SqliteDbPath, err)
		return false, err
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	if err = engine.Ping(); err != nil {
		rcali.Logger.Error("ping sqlitedb fail on ", SqliteDbPath, err)
		return false, err
	}

	if exist, err := engine.IsTableExist(&models.Author{}); !exist || err != nil {
		rcali.Logger.Error("table authors not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Book{}); !exist || err != nil {
		rcali.Logger.Error("table books not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.BookRatingLink{}); !exist || err != nil {
		rcali.Logger.Error("table books_ratings_link not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Comments{}); !exist || err != nil {
		rcali.Logger.Error("table comments not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Data{}); !exist || err != nil {
		rcali.Logger.Error("table data not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Feed{}); !exist || err != nil {
		rcali.Logger.Error("table feed not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Identifier{}); !exist || err != nil {
		rcali.Logger.Error("table identifies not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Language{}); !exist || err != nil {
		rcali.Logger.Error("table languages not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Publisher{}); !exist || err != nil {
		rcali.Logger.Error("table publishers not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Tag{}); !exist || err != nil {
		rcali.Logger.Error("table tags not exit", err)
		return false, err
	}

	//localengine
	userHome, _ := rcali.Home()
	if localEngine, err = xorm.NewEngine("sqlite3", path.Join(userHome, ".calilocaldb.db")); err != nil {
		rcali.Logger.Error("open sqlitedb fail on ", path.Join(userHome, ".calilocaldb.db"), err)
		return false, err
	}
	localEngine.ShowSQL(true)
	localEngine.Logger().SetLevel(core.LOG_DEBUG)
	if err = localEngine.Ping(); err != nil {
		rcali.Logger.Error("ping sqlitedb fail on ", path.Join(userHome, ".calilocaldb.db"), err)
		return false, err
	}

	//add user table
	//localEngine.CreateTables(new(models.UserInfo))
	if err = localEngine.Sync2(models.UserInfo{}); err != nil {
		return false, err
	}
	tmpInfo := models.UserInfo{}
	localEngine.ID("init").Get(&tmpInfo)
	if tmpInfo.Id != "init" {
		if _, err = localEngine.Insert(models.DefaultUserInfo); err != nil {
			return false, err
		}
	}
	tmpInfo = models.UserInfo{}
	localEngine.ID("admin").Get(&tmpInfo)
	if tmpInfo.Id != "admin" {
		if _, err = localEngine.Insert(models.DefaultAdminUserInfo); err != nil {
			return false, err
		}
	}

	//add role table
	if err = localEngine.Sync2(models.Role{}); err != nil {
		return false, err
	}
	roleInfo := models.Role{}
	localEngine.ID("admin").Get(&roleInfo)
	if roleInfo.Id != "admin" {
		_, err = localEngine.Insert(models.DefaultAdminRole)
		if err != nil {
			return false, err
		}
	}
	roleInfo = models.Role{}
	localEngine.ID("user").Get(&roleInfo)
	if roleInfo.Id != "user" {
		_, err = localEngine.Insert(models.DefaultUserRole)
		if err != nil {
			return false, err
		}
	}
	roleInfo = models.Role{}
	localEngine.ID("watcher").Get(&roleInfo)
	if roleInfo.Id != "watcher" {
		_, err = localEngine.Insert(models.DefaultWatcherRole)
		if err != nil {
			return false, err
		}
	}

	//add default user and role
	if err = localEngine.Sync2(models.UserInfoRoleLink{}); err != nil {
		return false, err
	}
	userRoleLinkInfo := models.UserInfoRoleLink{}
	localEngine.ID("user").Get(&userRoleLinkInfo)
	if userRoleLinkInfo.Id != "user" {
		if _, err = localEngine.Insert(models.DefaultUserInfoRole); err != nil {
			return false, err
		}
	}
	userRoleLinkInfo = models.UserInfoRoleLink{}
	localEngine.ID("admin").Get(&userRoleLinkInfo)
	if userRoleLinkInfo.Id != "admin" {
		if _, err = localEngine.Insert(models.DefaultAdminUserInfoRole); err != nil {
			return false, err
		}
	}

	//add role action
	roleAction := models.RoleAction{}
	err = localEngine.DropTables(roleAction)
	if err = localEngine.Sync2(models.RoleAction{}); err != nil {
		return false, err
	}
	if _, err = localEngine.Insert(models.RoleActions); err != nil {
		return false, err
	}

	//sysconfig add
	if err = localEngine.Sync2(models.SysConfig{}); err != nil {
		return false, err
	}
	for _, value := range models.DefaultSysConfig {
		sysConfig := models.SysConfig{}
		localEngine.ID(value.Id).Get(&sysConfig)
		if sysConfig.Id != value.Id {
			_, err = localEngine.Insert(value)
			if err != nil {
				return false, err
			}
		}
	}

	//sysstatus add
	if err = localEngine.Sync2(models.SysStatus{}); err != nil {
		return false, err
	}

	//sync user and book and confg
	if err = localEngine.Sync2(models.UserInfoBookUploadLink{}, models.UserInfoBookDownloadLink{}, models.UserConfig{}); err != nil {
		return false, err
	}

	//touch the metadb
	if _, err = localEngine.Exec("ATTACH DATABASE \"" + SqliteDbPath + "\" AS \"data\""); err != nil {
		return false, err
	}

	rcali.Logger.Info("----------DbInitOk----------")
	return true, nil

}
