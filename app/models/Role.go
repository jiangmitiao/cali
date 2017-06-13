package models

type Role struct {
	Id      string `json:"id" xorm:"pk 'id'"`
	Name    string `json:"name" xorm:"varchar(32) 'name'"`
	Comment string `json:"comment" xorm:"varchar(64) 'comment'"`
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
