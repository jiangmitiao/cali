package models

type Tag struct {
	Id   int    `json:"id" xorm:"pk 'id'"`
	Name string `json:"name" xorm:"'name'"`
}

func (Tag) TableName() string {
	return "tags"
}
