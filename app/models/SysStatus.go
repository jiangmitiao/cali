package models

type SysStatus struct {
	Id       string `json:"id" xorm:"pk varchar(64) 'id'"`
	Ikey     string `json:"key" xorm:"varchar(64) notnull 'ikey'"`
	Value    string `json:"value" xorm:"varchar(64) default '' 'value'"`
	Comments string `json:"comments" xorm:"varchar(128) default '' 'comments'"`

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (SysStatus) TableName() string {
	return "sys_status"
}
