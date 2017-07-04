package models

import "time"

type CaliCategory struct {
	Id       string `json:"id" xorm:"pk varchar(64) 'id'"`
	Category string `json:"category" xorm:"varchar(64) notnull 'category'"`

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (CaliCategory) TableName() string {
	return "cali_category"
}

var (
	DefaultCaliCategory = CaliCategory{Id: "default", Category: "全部", CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
)
