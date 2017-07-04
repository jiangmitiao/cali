package models

type SysConfig struct {
	Id       string `json:"id" xorm:"pk varchar(64) 'id'"`
	Ikey     string `json:"key" xorm:"varchar(64) notnull 'ikey'"`
	Value    string `json:"value" xorm:"varchar(64) notnull 'value'"`
	Comments string `json:"comments" xorm:"varchar(128) default '' 'comments'"`

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (SysConfig) TableName() string {
	return "sys_config"
}

var DefaultSysConfig = make([]SysConfig, 0)

func init() {
	DefaultSysConfig = append(DefaultSysConfig, SysConfig{Id: "cnzzid", Ikey: "cnzzid", Value: "1262308688", Comments: "cnzz data monitor"})

	DefaultSysConfig = append(DefaultSysConfig, SysConfig{Id: "openregist", Ikey: "openregist", Value: "true", Comments: "open rigist?"})

	DefaultSysConfig = append(DefaultSysConfig, SysConfig{Id: "alldownloadlimit", Ikey: "alldownloadlimit", Value: "10000", Comments: "how many books can download one day?"})

	//iplimit
	DefaultSysConfig = append(DefaultSysConfig, SysConfig{Id: "iplimit", Ikey: "iplimit", Value: "60", Comments: ""})
}
