package models

import "time"

type Book struct {
	Id           int       `json:"id" xorm:"pk 'id'"`
	Title        string    `json:"title" xorm:"'title'"`
	Sort         string    `json:"sort" xorm:"'sort'"`
	Timestamp    time.Time `json:"timestamp"  xorm:"'timestamp'"`
	Pubdate      time.Time `json:"pubdate"  xorm:"'pubdate'"`
	SeriesIndex  float64   `json:"series_index"  xorm:"'series_index'"`
	AuthorSort   string    `json:"author_sort"  xorm:"'author_sort'"`
	Isbn         string    `json:"isbn"  xorm:"'isbn'"`
	Lccn         string    `json:"lccn"  xorm:"'lccn'"`
	Path         string    `json:"path"  xorm:"'path'"`
	Flags        int       `json:"flags"  xorm:"'flags'"`
	Uuid         string    `json:"uuid"  xorm:"'uuid'"`
	HasCover     int       `json:"has_cover"  xorm:"'has_cover'"`
	LastModified time.Time `json:"last_modified"  xorm:"'last_modified'"`
}

func (Book) TableName() string {
	return "books"
}

type BookVo struct {
	Id           int       `json:"id" xorm:"pk 'id'"`
	Title        string    `json:"title" xorm:"'title'"` // book name
	Sort         string    `json:"sort" xorm:"'sort'"`
	Timestamp    time.Time `json:"timestamp"  xorm:"'timestamp'"`
	Pubdate      time.Time `json:"pubdate"  xorm:"'pubdate'"`
	SeriesIndex  float64   `json:"series_index"  xorm:"'series_index'"`
	AuthorSort   string    `json:"author_sort"  xorm:"'author_sort'"`
	Isbn         string    `json:"isbn"  xorm:"'isbn'"`
	Lccn         string    `json:"lccn"  xorm:"'lccn'"`
	Path         string    `json:"path"  xorm:"'path'"`
	Flags        int       `json:"flags"  xorm:"'flags'"`
	Uuid         string    `json:"uuid"  xorm:"'uuid'"`
	HasCover     int       `json:"has_cover"  xorm:"'has_cover'"`
	LastModified time.Time `json:"last_modified"  xorm:"'last_modified'"`

	Rating int `json:"rating" xorm:"'rating'"`

	Name string `json:"name" xorm:"'name'"` //author name

	Comments string `json:"comments" xorm:"'comments'"` //comments's text only use in single book

	Format string `json:"format" xorm:"'format'"`
}
