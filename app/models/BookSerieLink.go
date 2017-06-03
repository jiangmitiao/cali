package models

type BookSerieLink struct {
	Id     int `json:"id" xorm:"pk 'id'"`
	Book   int `json:"book" xorm:"'book'"`
	Series int `json:"series" xorm:"'series'"`
}

func (BookSerieLink) TableName() string {
	return "books_series_link"
}
