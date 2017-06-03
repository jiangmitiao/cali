package models

type Publisher struct {
	Id   int    `json:"id" xorm:"pk 'id'"`
	Name string `json:"name" xorm:"'name'"`
	Sort string `json:"sort" sort:"'sort'"`
}

func (Publisher) TableName() string {
	return "publishers"
}
