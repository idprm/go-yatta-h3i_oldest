package models

type Keyword struct {
	Id         *int `gorm:"primary_key"`
	Name       *string
	Adnet      *string
	IsPostBack *bool
}

func (Keyword) TableName() string {
	return "keywords"
}
