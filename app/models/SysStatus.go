package models

type SysStatus struct {
	Id       string `json:"id" xorm:"pk 'id'"`
	Key      string `json:"key" xorm:"'key'"`
	Value    string `json:"value" xorm:"'value'"`
	Comments string `json:"comments" xorm:"'comments'"`
}

func (SysStatus) TableName() string {
	return "sys_status"
}
