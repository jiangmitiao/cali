package models

type ConversionOption struct {
	Id     int    `json:"id" xorm:"pk 'id'"`
	Format string `json:"format" xorm:"'format'"`
	Book   int    `json:"book" xorm:"'book'"`
	Data   string `json:"data" xorm:"'data'"`
}

func (ConversionOption) TableName() string {
	return "conversion_options"
}
