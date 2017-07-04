package models

type CaliBook struct {
	Id     string `json:"id" xorm:"pk varchar(64) 'id'"`
	Title  string `json:"title" xorm:"varchar(64) not null 'title'"`
	Author string `json:"author" xorm:"varchar(64) not null 'author'"`

	DownloadCount int    `json:"download_count" xorm:"int default 0 'download_count'"`
	DoubanId      string `json:"douban_id" xorm:"varchar(16) default '' 'douban_id'"`
	DoubanJson    string `json:"douban_json" xorm:"text 'douban_json'"`

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (CaliBook) TableName() string {
	return "cali_book"
}

type CaliBookVo struct {
	CaliBook

	Formats []CaliFormat `json:"formats"`

	Categories []CaliCategory `json:"categories"`
}
