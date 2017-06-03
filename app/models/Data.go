package models

type Data struct {
	Id               int    `json:"id" xorm:"pk 'id'"`
	Book             int    `json:"book" xorm:"'book'"`
	Format           string `json:"format" xorm:"'format'"`
	UncompressedSize int    `json:"uncompressed_size" xorm:"'uncompressed_size'"`
	Name             string `json:"name" xorm:"'name'"`
}

func (Data) TableName() string {
	return "data"
}
