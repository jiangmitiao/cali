package models

type CaliBookCategory struct {
	Id           string `json:"id" xorm:"pk varchar(64) 'id'"`
	CaliCategory string `json:"cali_category" xorm:"varchar(64) notnull 'cali_category'"`
	CaliBook     string `json:"cali_book" xorm:"varchar(64) notnull 'cali_book'"`

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (CaliBookCategory) TableName() string {
	return "cali_book_category"
}
