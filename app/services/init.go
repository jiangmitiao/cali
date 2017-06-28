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

	DefaultBookService       = CaliBookService{lock: &sync.Mutex{}}
	DefaultFormatService	 = CaliFormatService{}
	DefaultUserService       = UserService{}
	DefaultUserRoleService   = UserRoleService{}
	DefaultRoleActionService = RoleActionService{}
	DefaultSysConfigService  = SysConfigService{}
	DefaultSysStatusService  = SysStatusService{lock: &sync.Mutex{}}
	DefaultUserConfigService = UserConfigService{}
)

//init the db,should take a db filepath
func DbInit(SqliteDbPath string) (bool, error) { //username, password, host, database string
	SqliteDbPath =  path.Join(SqliteDbPath,"cali.db")
	//OPEN
	var err error
	if engine, err = xorm.NewEngine("sqlite3", SqliteDbPath); err != nil {
		rcali.Logger.Error("open sqlitedb fail on ", SqliteDbPath, err)
		return false, err
	}

	//CONFIG CHECK
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	if err = engine.Ping(); err != nil {
		rcali.Logger.Error("ping sqlitedb fail on ", SqliteDbPath, err)
		return false, err
	}

	//BOOKS
	if err = engine.Sync2(models.CaliBook{},models.CaliFormat{},models.CaliCategory{},models.CaliBookCategory{}); err != nil {
		return false, err
	}

	//USERS
	if err = engine.Sync2(models.UserInfo{}); err != nil {
		return false, err
	}
	tmpInfo := models.UserInfo{}
	engine.ID("init").Get(&tmpInfo)
	if tmpInfo.Id != "init" {
		if _, err = engine.Insert(models.DefaultUserInfo); err != nil {
			return false, err
		}
	}
	tmpInfo = models.UserInfo{}
	engine.ID("admin").Get(&tmpInfo)
	if tmpInfo.Id != "admin" {
		if _, err = engine.Insert(models.DefaultAdminUserInfo); err != nil {
			return false, err
		}
	}

	//add role table
	if err = engine.Sync2(models.Role{}); err != nil {
		return false, err
	}
	roleInfo := models.Role{}
	engine.ID("admin").Get(&roleInfo)
	if roleInfo.Id != "admin" {
		_, err = engine.Insert(models.DefaultAdminRole)
		if err != nil {
			return false, err
		}
	}
	roleInfo = models.Role{}
	engine.ID("user").Get(&roleInfo)
	if roleInfo.Id != "user" {
		_, err = engine.Insert(models.DefaultUserRole)
		if err != nil {
			return false, err
		}
	}
	roleInfo = models.Role{}
	engine.ID("watcher").Get(&roleInfo)
	if roleInfo.Id != "watcher" {
		_, err = engine.Insert(models.DefaultWatcherRole)
		if err != nil {
			return false, err
		}
	}

	//add default user and role
	if err = engine.Sync2(models.UserInfoRoleLink{}); err != nil {
		return false, err
	}
	userRoleLinkInfo := models.UserInfoRoleLink{}
	engine.ID("user").Get(&userRoleLinkInfo)
	if userRoleLinkInfo.Id != "user" {
		if _, err = engine.Insert(models.DefaultUserInfoRole); err != nil {
			return false, err
		}
	}
	userRoleLinkInfo = models.UserInfoRoleLink{}
	engine.ID("admin").Get(&userRoleLinkInfo)
	if userRoleLinkInfo.Id != "admin" {
		if _, err = engine.Insert(models.DefaultAdminUserInfoRole); err != nil {
			return false, err
		}
	}

	//add role action
	roleAction := models.RoleAction{}
	err = engine.DropTables(roleAction)
	if err = engine.Sync2(models.RoleAction{}); err != nil {
		return false, err
	}
	if _, err = engine.Insert(models.RoleActions); err != nil {
		return false, err
	}

	//sysconfig add
	if err = engine.Sync2(models.SysConfig{}); err != nil {
		return false, err
	}
	for _, value := range models.DefaultSysConfig {
		sysConfig := models.SysConfig{}
		engine.ID(value.Id).Get(&sysConfig)
		if sysConfig.Id != value.Id {
			_, err = engine.Insert(value)
			if err != nil {
				return false, err
			}
		}
	}

	//sysstatus add
	if err = engine.Sync2(models.SysStatus{}); err != nil {
		return false, err
	}

	//sync user and book and confg
	if err = engine.Sync2(models.UserInfoBookUploadLink{}, models.UserInfoBookDownloadLink{}, models.UserConfig{}); err != nil {
		return false, err
	}

	rcali.Logger.Info("----------DbInitOk----------")
	return true, nil

}
