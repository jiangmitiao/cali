package models

import "time"

type CaliCategory struct {
	Id       string `json:"id" xorm:"pk 'id'"`
	Category string `json:"category" xorm:"'category'"`

	CreatedAt int64 `json:"created" xorm:"'created'"`
	UpdatedAt int64 `json:"updated" xorm:"'updated'"`
}

func (CaliCategory) TableName() string {
	return "cali_category"
}


var (
	DefaultCaliCategory = CaliCategory{Id:"default",Category:"默认",CreatedAt:time.Now().Unix(),UpdatedAt:time.Now().Unix()}
)