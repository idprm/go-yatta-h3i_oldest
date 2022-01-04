package models

import (
	"time"
)

type Subscription struct {
	Id               uint64 `gorm:"primary_key"`
	Msisdn           int
	Service          string
	Adn              string
	Operator         string
	Keyword          string
	SubKeywork       string
	ChannelSubscribe string
	IsActive         bool
	MsgTimestamp     time.Time
	RenewalAt        time.Time
	ConfirmAt        time.Time
	PurgeAt          time.Time
}

func (Subscription) TableName() string {
	return "subscription"
}
