package models

type SysConfig struct {
	Id       string `json:"id" xorm:"pk 'id'"`
	Key      string `json:"key" xorm:"'key'"`
	Value    string `json:"value" xorm:"'value'"`
	Comments string `json:"comments" xorm:"'comments'"`
}

func (SysConfig) TableName() string {
	return "sys_config"
}

var DefaultSysConfig = make([]SysConfig, 0)

func init() {
	DefaultSysConfig = append(DefaultSysConfig, SysConfig{Id: "cnzzid", Key: "cnzzid", Value: "1262308688", Comments: "cnzz data monitor"})

	DefaultSysConfig = append(DefaultSysConfig, SysConfig{Id: "openregist", Key: "openregist", Value: "true", Comments: "open rigist?"})

	DefaultSysConfig = append(DefaultSysConfig, SysConfig{Id: "alldownloadlimit", Key: "alldownloadlimit", Value: "10000", Comments: "how many books can download one day?"})
}
