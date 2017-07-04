package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

type UserInfoNew struct {
	Id            string `json:"id" xorm:"pk 'id'"`
	LoginEmail    string `json:"loginEmail" xorm:"varchar(64) notnull default 'error' 'login_email'"`
	LoginPassword string `json:"loginPassword" xorm:"varchar(128) notnull 'login_password'"`
	Salt          string `json:"salt" xorm:"varchar(128) notnull 'salt'"`
	Email         string `json:"email" xorm:"varchar(128) 'email'"`

	UserName string `json:"userName" xorm:"varchar(64) notnull 'user_name'"`
	Img      string `json:"img" xorm:"varchar(256) 'img'"`

	Valid int `json:"valid" xorm:"int default 0 'valid'"` //0 有效 1 无效  2 wait active

	CreatedAt int64 `json:"created" xorm:"'created'"`
	UpdatedAt int64 `json:"updated" xorm:"'updated'"`
}

func (UserInfoNew) TableName() string {
	return "user_info"
}

func TestModel(t *testing.T) {
	var engine *xorm.Engine
	var err error
	if engine, err = xorm.NewEngine("sqlite3", "./test.db"); err != nil {
		panic(err)
	}

	//CONFIG CHECK
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	if err = engine.Ping(); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", engine.Sync2(UserInfo{}))
	fmt.Printf("%#v\n", engine.Sync2(UserInfoNew{}))

	engine.Close()
	os.Remove("./test.db")
}

func TestModelNew(t *testing.T) {
	var engine *xorm.Engine
	var err error
	if engine, err = xorm.NewEngine("mysql", "cali:calipassword@tcp(127.0.0.1:3306)/cali?charset=utf8"); err != nil {
		fmt.Println(err)
		return
	}

	//CONFIG CHECK
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	if err = engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	engine.DropTables(UserInfo{}, CaliBook{}, CaliBookCategory{},
		CaliCategory{}, CaliFormat{}, Role{},
		RoleAction{}, SysConfig{}, SysStatus{},
		UserConfig{}, UserInfoBookDownloadLink{}, UserInfoBookUploadLink{}, UserInfoRoleLink{},
	)

	fmt.Printf("%#v\n", engine.Sync2(UserInfo{}))
	fmt.Printf("%#v\n", engine.Sync2(CaliBook{}))
	fmt.Printf("%#v\n", engine.Sync2(CaliBookCategory{}))
	fmt.Printf("%#v\n", engine.Sync2(CaliCategory{}))
	fmt.Printf("%#v\n", engine.Sync2(CaliFormat{}))
	fmt.Printf("%#v\n", engine.Sync2(Role{}))
	fmt.Printf("%#v\n", engine.Sync2(RoleAction{}))
	fmt.Printf("%#v\n", engine.Sync2(SysConfig{}))
	fmt.Printf("%#v\n", engine.Sync2(SysStatus{}))
	fmt.Printf("%#v\n", engine.Sync2(UserConfig{}))
	fmt.Printf("%#v\n", engine.Sync2(UserInfoBookDownloadLink{}))
	fmt.Printf("%#v\n", engine.Sync2(UserInfoBookUploadLink{}))
	fmt.Printf("%#v\n", engine.Sync2(UserInfoRoleLink{}))

	//fmt.Printf("%#v\n",engine.Sync2(UserInfoNew{}))

	//fmt.Printf("%#v\n",engine.DropTables)

	engine.Close()
}
