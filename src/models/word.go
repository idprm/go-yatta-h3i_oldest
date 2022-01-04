package models

type Word struct {
	Id   int `gorm:"primary_key"`
	Slug string
	Name string
}

func (Word) TableName() string {
	return "words"
}
