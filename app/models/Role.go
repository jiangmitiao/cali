package models

type Role struct {
	Id      string `json:"id" xorm:"pk varchar(64) 'id'"`
	Name    string `json:"name" xorm:"varchar(32) notnull 'name'"`
	Comment string `json:"comment" xorm:"varchar(64) notnull 'comment'"`
}

func (Role) TableName() string {
	return "role"
}

var DefaultAdminRole = Role{
	Id:      "admin",
	Name:    "admin",
	Comment: "admin",
}

var DefaultUserRole = Role{
	Id:      "user",
	Name:    "user",
	Comment: "user",
}

var DefaultWatcherRole = Role{
	Id:      "watcher",
	Name:    "watcher",
	Comment: "watcher",
}
