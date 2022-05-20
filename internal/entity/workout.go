package entity

import "time"

type Workout struct {
	Id          uint
	PaymentDate time.Time
	EndDate     time.Time
	Remains     uint8
}
