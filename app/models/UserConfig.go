package models

import "github.com/google/uuid"

type UserConfig struct {
	Id          string `json:"id" xorm:"pk 'id'"`
	UserInfo    string `json:"user_info" xorm:"'user_info'"`
	MaxDownload int    `json:"max_download" xorm:"'max_download' default 1"`
}

func (UserConfig) TableName() string {
	return "user_config"
}

func NewUserConfig(userInfoId string) UserConfig {
	return UserConfig{Id: uuid.New().String(), UserInfo: userInfoId, MaxDownload: 1}
}
