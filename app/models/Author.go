package models

type Author struct {
	Id   int    `json:"id" xorm:"pk 'id'"`
	Name string `json:"name" xorm:"'name'"`
	Sort string `json:"sort" xorm:"'sort'"`
	Link string `json:"link" xorm:"'link'"`
}

func (Author) TableName() string {
	return "authors"
}
