package models

type CaliCategory struct {
	Id string `json:"id" xorm:"pk 'id'"`
	Category string `json:"category" xorm:"'category'"`
}

func (CaliCategory) TableName() string {
	return "cali_category"
}