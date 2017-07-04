package models

type CaliFormat struct {
	Id       string `json:"id" xorm:"pk varchar(64) 'id'"`
	CaliBook string `json:"cali_book" xorm:"varchar(64) notnull 'cali_book'"`

	Format           string `json:"format" xorm:"varchar(64) notnull 'format'"`
	UncompressedSize int64  `json:"uncompressed_size" xorm:"bigint default 0 'uncompressed_size'"`
	Title            string `json:"title" xorm:"varchar(64) notnull 'title'"`
	Author           string `json:"author" xorm:"varchar(64) notnull 'author'"`
	FileName         string `json:"file_name" xorm:"varchar(128) notnull 'file_name'"`
	Tag              string `json:"tag" xorm:"varchar(128) default '' 'tag'"`

	DownloadCount int `json:"download_count" xorm:"int default 0 'download_count'"`

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (CaliFormat) TableName() string {
	return "cali_format"
}
