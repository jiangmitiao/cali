package models

type Comments struct {
	Id   int    `json:"id" xorm:"pk 'id'"`
	Book int    `json:"book" xorm:"'book'"`
	Text string `json:"text" xorm:"'text'"`
}

func (Comments) TableName() string {
	return "comments"
}
