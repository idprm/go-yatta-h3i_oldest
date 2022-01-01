package models

import "time"

type Subscription struct {
	Id               uint64
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
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
