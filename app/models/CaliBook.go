package models


type CaliBook struct {
	Id string `json:"id" xorm:"pk 'id'"`
	Title string `json:"title" xorm:"'title'"`
	Author string `json:"author" xorm:"'author'"`

	DownloadCount int `json:"download_count" xorm:"'download_count'"`
	DoubanId string  `json:"douban_id" xorm:"'douban_id'"`
}

func (CaliBook) TableName() string {
	return "cali_book"
}

type CaliBookVo struct {
	CaliBook

	Formats []CaliFormat

	Categories  []CaliCategory
}