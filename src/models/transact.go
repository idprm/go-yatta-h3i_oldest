package models

import "time"

type Transact struct {
	Id           uint64
	TrxId        uint64
	Name         string
	Msisdn       string
	MsgIndex     string
	MsgTimestamp time.Time
}
