package models

type UserInfoRoleLink struct {
	Id       string `json:"id" xorm:"pk varchar(64) 'id'"`
	UserInfo string `json:"user_info" xorm:"varchar(64) notnull 'user_info'"`
	Role     string `json:"role" xorm:"varchar(64) notnull 'role'"`
}

func (UserInfoRoleLink) TableName() string {
	return "user_info_role_link"
}

var DefaultUserInfoRole = UserInfoRoleLink{
	Id:       "user",
	UserInfo: "init",
	Role:     "user",
}

var DefaultAdminUserInfoRole = UserInfoRoleLink{
	Id:       "admin",
	UserInfo: "admin",
	Role:     "admin",
}
