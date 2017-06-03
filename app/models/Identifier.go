package models

type Identifier struct {
	Id   int    `json:"id" xorm:"pk 'id'"`
	Book int    `json:"book" xorm:"'book'"`
	Type string `json:"type"  xorm:"'type'"`
	Val  string `json:"val"  xorm:"'val'"`
}

func (Identifier) TableName() string {
	return "identifiers"
}
