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
	engine *xorm.Engine

	DefaultBookService       = CaliBookService{lock: &sync.Mutex{}}
	DefaultFormatService     = CaliFormatService{}
	DefaultCategoryService   = CaliCategoryService{}
	DefaultUserService       = UserService{}
	DefaultUserRoleService   = UserRoleService{}
	DefaultRoleActionService = RoleActionService{}
	DefaultSysConfigService  = SysConfigService{}
	DefaultSysStatusService  = SysStatusService{lock: &sync.Mutex{}}
	DefaultUserConfigService = UserConfigService{}
)

func DbInitBySqlite(sqliteDbPath string) error {
	sqliteDbPath = path.Join(sqliteDbPath, "cali.db")
	var err error
	if engine, err = xorm.NewEngine("sqlite3", sqliteDbPath); err != nil {
		rcali.Logger.Error("open sqlitedb fail on ", sqliteDbPath, err)
		return err
	}
	return nil
}

func DbInitByMysql(mysqlDsn string) error {
	var err error
	if engine, err = xorm.NewEngine("mysql", mysqlDsn); err != nil {
		rcali.Logger.Error("open mysql fail on ", mysqlDsn, err)
		return err
	}
	return nil
}

//init the db,should take a db filepath
func DbInit() (bool, error) { //username, password, host, database string
	var err error
	//CONFIG CHECK
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	if err = engine.Ping(); err != nil {
		rcali.Logger.Error("ping sqlitedb fail on ", err)
		return false, err
	}

	//BOOKS
	if err = engine.Sync2(models.CaliBook{}, models.CaliFormat{}, models.CaliCategory{}, models.CaliBookCategory{}); err != nil {
		return false, err
	}
	engine.InsertOne(models.DefaultCaliCategory)

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

func UpdateSqlite2Mysql()  {
	if dbPath, dbPathFund := rcali.GetSqliteDbPath(); dbPathFund {
		dbPath = path.Join(dbPath, "cali.db")
		if has,_ :=rcali.FileExists(dbPath);has {
			if srcEngine, err := xorm.NewEngine("sqlite3", dbPath); err == nil {
				UpdateSql2Sql(engine,srcEngine)
				defer srcEngine.Close()
			}
		}
	}
}

func UpdateSql2Sql(dst *xorm.Engine, src *xorm.Engine) {
	if has,_:=src.IsTableExist(models.CaliBook{});has {
		books := make([]models.CaliBook, 0)
		src.Find(&books)
		if affected, err := dst.Insert(books); err == nil && int(affected) == len(books) {
			src.Delete(models.CaliBook{})
			src.DropTables(models.CaliBook{})
		}
	}

	if has,_:=src.IsTableExist(models.CaliCategory{});has {
		categories := make([]models.CaliCategory, 0)
		src.Find(&categories)
		if affected, err := dst.Insert(categories); err == nil && int(affected) == len(categories) {
			src.Delete(models.CaliCategory{})
			src.DropTables(models.CaliCategory{})
		}
	}

	if has,_:=src.IsTableExist(models.CaliBookCategory{});has {
		bookCategories := make([]models.CaliBookCategory, 0)
		src.Find(&bookCategories)
		if affected, err := dst.Insert(bookCategories); err == nil && int(affected) == len(bookCategories) {
			src.Delete(models.CaliBookCategory{})
			src.DropTables(models.CaliBookCategory{})
		}
	}

	if has,_:=src.IsTableExist(models.CaliFormat{});has {
		formats := make([]models.CaliFormat, 0)
		src.Find(&formats)
		if affected, err := dst.Insert(formats); err == nil && int(affected) == len(formats) {
			src.Delete(models.CaliFormat{})
			src.DropTables(models.CaliFormat{})
		}
	}

	if has,_:=src.IsTableExist(models.UserConfig{});has {
		userConfigs := make([]models.UserConfig, 0)
		src.Find(&userConfigs)
		if affected, err := dst.Insert(userConfigs); err == nil && int(affected) == len(userConfigs) {
			src.Delete(models.UserConfig{})
			src.DropTables(models.UserConfig{})
		}
	}

	if has,_:=src.IsTableExist(models.UserInfo{});has {
		users := make([]models.UserInfo, 0)
		src.Find(&users)
		if affected, err := dst.Insert(users); err == nil && int(affected) == len(users) {
			src.Delete(models.UserInfo{})
			src.DropTables(models.UserInfo{})
		}
	}

	if has,_:=src.IsTableExist(models.UserInfoBookDownloadLink{});has {
		userDowns := make([]models.UserInfoBookDownloadLink, 0)
		src.Find(&userDowns)
		if affected, err := dst.Insert(userDowns); err == nil && int(affected) == len(userDowns) {
			src.Delete(models.UserInfoBookDownloadLink{})
			src.DropTables(models.UserInfoBookDownloadLink{})
		}
	}

	if has,_:=src.IsTableExist(models.UserInfoBookUploadLink{});has {
		userUps := make([]models.UserInfoBookUploadLink, 0)
		src.Find(&userUps)
		if affected, err := dst.Insert(userUps); err == nil && int(affected) == len(userUps) {
			src.Delete(models.UserInfoBookUploadLink{})
			src.DropTables(models.UserInfoBookUploadLink{})
		}
	}

	if has,_:=src.IsTableExist(models.UserInfoRoleLink{});has {
		userRoles := make([]models.UserInfoRoleLink, 0)
		src.Find(&userRoles)
		if affected, err := dst.Insert(userRoles); err == nil && int(affected) == len(userRoles) {
			src.Delete(models.UserInfoRoleLink{})
			src.DropTables(models.UserInfoRoleLink{})
		}
	}

}
