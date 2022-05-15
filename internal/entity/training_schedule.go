package entity

import "time"

type TrainingSchedule struct {
	PaymentDate time.Time
	EndDate     time.Time
	Remains     uint8
}
