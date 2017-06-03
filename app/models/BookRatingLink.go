package models

type BookRatingLink struct {
	Id     int `json:"id" xorm:"pk 'id'"`
	Book   int `json:"book" xorm:"'book'"`
	Rating int `json:"rating" xorm:"'rating'"`
}

func (BookRatingLink) TableName() string {
	return "books_ratings_link"
}
