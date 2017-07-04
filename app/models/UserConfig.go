package models

import "github.com/google/uuid"

type UserConfig struct {
	Id             string `json:"id" xorm:"pk varchar(64) 'id'"`
	UserInfo       string `json:"user_info" xorm:"varchar(64) notnull 'user_info'"`
	MaxDownload    int    `json:"max_download" xorm:"int default 5 'max_download'"`
	TmpMaxDownload int    `json:"tmpMaxDownload" xorm:"int default 0 'tmp_max_download'"`
}

func (UserConfig) TableName() string {
	return "user_config"
}

func NewUserConfig(userInfoId string) UserConfig {
	return UserConfig{Id: uuid.New().String(), UserInfo: userInfoId, MaxDownload: 5}
}
