package models

type UserInfoBookUploadLink struct {
	Id        string `json:"id" xorm:"pk 'id'"`
	UserInfo  string `json:"user_info" xorm:"'user_info'"`
	Book      int    `json:"book" xorm:"'book'"`
	CreatedAt int64  `json:"created" xorm:"created"`
	UpdatedAt int64  `json:"updated" xorm:"updated"`
}

func (UserInfoBookUploadLink) TableName() string {
	return "user_info_book_upload_link"
}
