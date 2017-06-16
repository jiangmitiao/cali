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
}
