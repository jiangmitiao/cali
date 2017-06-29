package models

type UserInfoBookDownloadLink struct {
	Id         string `json:"id" xorm:"pk 'id'"`
	UserInfo   string `json:"user_info" xorm:"'user_info'"`
	CaliFormat string `json:"cali_format" xorm:"'cali_format'"`
	CreatedAt  int64  `json:"created" xorm:"'created'"`
	UpdatedAt  int64  `json:"updated" xorm:"'updated'"`
}

func (UserInfoBookDownloadLink) TableName() string {
	return "user_info_book_download_link"
}
