package models

type CaliFormat struct {
	Id string `json:"id" xorm:"pk 'id'"`
	CaliBook string `json:"cali_book" xorm:"'cali_book'"`

	Format string `json:"format" xorm:"'format'"`
	UncompressedSize int64 `json:"uncompressed_size" xorm:"'uncompressed_size'"`
	FileName string `json:"file_name" xorm:"'file_name'"`
	Tag string `json:"tag" xorm:"'tag'"`
}

func (CaliFormat) TableName() string {
	return "cali_format"
}
