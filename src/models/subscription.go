package models

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
	SubscribedFrom   []byte
	SubscribedUntil  []byte
}
