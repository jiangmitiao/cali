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
)

var engine *xorm.Engine

var localEngine *xorm.Engine

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//init the db,should take a db filepath
func DbInit(SqliteDbPath string) (bool, error) { //username, password, host, database string
	if bool, err := rcali.FileExists(SqliteDbPath); !bool {
		rcali.DEBUG.Debug("sqlitedbpath is error", SqliteDbPath, err)
		return false, err
	}

	var err error
	engine, err = xorm.NewEngine("sqlite3", SqliteDbPath)
	if err != nil {
		rcali.DEBUG.Debug("open sqlitedb fail on ", SqliteDbPath, err)
		return false, err
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	err = engine.Ping()
	if err != nil {
		rcali.DEBUG.Debug("ping sqlitedb fail on ", SqliteDbPath, err)
		return false, err
	}

	if exist, err := engine.IsTableExist(&models.Author{}); !exist || err != nil {
		rcali.DEBUG.Debug("table authors not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Book{}); !exist || err != nil {
		rcali.DEBUG.Debug("table books not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.BookRatingLink{}); !exist || err != nil {
		rcali.DEBUG.Debug("table books_ratings_link not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Comments{}); !exist || err != nil {
		rcali.DEBUG.Debug("table comments not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Data{}); !exist || err != nil {
		rcali.DEBUG.Debug("table data not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Feed{}); !exist || err != nil {
		rcali.DEBUG.Debug("table feed not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Identifier{}); !exist || err != nil {
		rcali.DEBUG.Debug("table identifies not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Language{}); !exist || err != nil {
		rcali.DEBUG.Debug("table languages not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Publisher{}); !exist || err != nil {
		rcali.DEBUG.Debug("table publishers not exit", err)
		return false, err
	}
	if exist, err := engine.IsTableExist(&models.Tag{}); !exist || err != nil {
		rcali.DEBUG.Debug("table tags not exit", err)
		return false, err
	}

	//localengine
	userHome, _ := rcali.Home()
	localEngine, err = xorm.NewEngine("sqlite3", path.Join(userHome, ".calilocaldb.db"))
	if err != nil {
		rcali.DEBUG.Debug("open sqlitedb fail on ", path.Join(userHome, ".calilocaldb.db"), err)
		return false, err
	}
	localEngine.ShowSQL(true)
	localEngine.Logger().SetLevel(core.LOG_DEBUG)
	err = localEngine.Ping()
	if err != nil {
		rcali.DEBUG.Debug("ping sqlitedb fail on ", path.Join(userHome, ".calilocaldb.db"), err)
		return false, err
	}
	//add user table
	if exist, err := localEngine.IsTableExist(&models.UserInfo{}); !exist || err != nil {
		//localEngine.CreateTables(new(models.UserInfo))
		err = localEngine.Sync2(models.UserInfo{})
		panic(err)
	}
	tmpInfo := models.UserInfo{}
	localEngine.ID("init").Get(&tmpInfo)
	if tmpInfo.Id != "init" {
		_, err = localEngine.Insert(models.DefaultUserInfo)
		panic(err)
	}

	rcali.DEBUG.Debug("----------DbInitOk----------")
	return true, nil

}
