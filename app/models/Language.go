package models

type Language struct {
	Id       int    `json:"id" xorm:"pk 'id'"`
	LangCode string `json:"lang_code" xorm:"'lang_code'"`
}

func (Language) TableName() string {
	return "languages"
}
