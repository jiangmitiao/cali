package models

type Feed struct {
	Id     int    `json:"id" xorm:"pk 'id'"`
	Title  string `json:"title" xorm:"'title'"`
	Script string `json:"script" xorm:"'script'"`
}

func (Feed) TableName() string {
	return "feeds"
}
