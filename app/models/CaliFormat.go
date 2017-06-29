package models

type CaliFormat struct {
	Id       string `json:"id" xorm:"pk 'id'"`
	CaliBook string `json:"cali_book" xorm:"'cali_book'"`

	Format           string `json:"format" xorm:"'format'"`
	UncompressedSize int64  `json:"uncompressed_size" xorm:"'uncompressed_size'"`
	Title            string `json:"title" xorm:"'title'"`
	Author           string `json:"author" xorm:"'author'"`
	FileName         string `json:"file_name" xorm:"'file_name'"`
	Tag              string `json:"tag" xorm:"'tag'"`

	DownloadCount int `json:"download_count" xorm:"'download_count'"`

	CreatedAt int64 `json:"created" xorm:"'created'"`
	UpdatedAt int64 `json:"updated" xorm:"'updated'"`
}

func (CaliFormat) TableName() string {
	return "cali_format"
}
