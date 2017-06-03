package models

type CustomColumn struct {
	Id            int `json:"id" xorm:"pk 'id'"`
	Label         string
	Name          string
	Datatype      string
	MarkForDelete int
	Editable      int
	Display       string
	IsMultiple    int
	Normalized    int
}

func (CustomColumn) TableName() string {
	return "custom_columns"
}
