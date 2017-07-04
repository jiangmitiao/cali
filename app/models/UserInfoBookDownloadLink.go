package models

type UserInfoBookDownloadLink struct {
	Id         string `json:"id" xorm:"pk varchar(64) 'id'"`
	UserInfo   string `json:"user_info" xorm:"varchar(64) notnull 'user_info'"`
	CaliFormat string `json:"cali_format" xorm:"varchar(64) notnull 'cali_format'"`

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (UserInfoBookDownloadLink) TableName() string {
	return "user_info_book_download_link"
}
