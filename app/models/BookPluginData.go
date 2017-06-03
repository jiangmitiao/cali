package models

type BookPluginData struct {
	Id   int    `json:"id" xorm:"pk 'id'"`
	Book int    `json:"book" xorm:"'book'"`
	Name string `json:"name"  xorm:"'name'"`
	Val  string `json:"val"  xorm:"'val'"`
}

func (BookPluginData) TableName() string {
	return "books_plugin_data"
}
