package models

import "time"

type Transact struct {
	Id           uint64    `gorm:"ID"`
	Msisdn       int       `gorm:"MSISDN"`
	MsgTimestamp string    `gorm:"MSGTIMESTAMP"`
	Adn          string    `gorm:"ADN"`
	OperatorId   string    `gorm:"OPERATORID"`
	Keyword      string    `gorm:"KEYWORD"`
	SubKeyword   string    `gorm:"SUBKEYWORD"`
	MsgData      string    `gorm:"MSGDATA"`
	MsgStatus    string    `gorm:"MSGSTATUS"`
	CloseReason  string    `gorm:"CLOSEREASON"`
	Channel      string    `gorm:"CHANNEL"`
	Service      string    `gorm:"SERVICE"`
	Price        string    `gorm:"PRICE"`
	SubmitedId   string    `gorm:"SUBMITEDID"`
	CreatedAt    time.Time `gorm:"CREATED_AT"`
	UpdatedAt    time.Time `gorm:"UPDATED_AT"`
}
