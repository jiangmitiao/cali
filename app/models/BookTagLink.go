package models

type BookTagLink struct {
	Id   int `json:"id" xorm:"pk 'id'"`
	Book int `json:"book" xorm:"'book'"`
	Tag  int `json:"series" xorm:"'tag'"`
}

func (BookTagLink) TableName() string {
	return "book_tag_link"
}
