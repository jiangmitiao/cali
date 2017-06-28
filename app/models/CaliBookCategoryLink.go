package models

type CaliBookCategory struct {
	Id string `json:"id" xorm:"pk 'id'"`
	CaliCategory string `json:"cali_category" xorm:"'cali_category'"`
	CaliBook string  `json:"cali_book" xorm:"'cali_book'"`
}

func (CaliBookCategory) TableName() string {
	return "cali_book_category"
}
