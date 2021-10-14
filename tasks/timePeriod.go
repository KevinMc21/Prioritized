package tasks

import "time"

type timePeriod struct{
	TimeStart	time.Time	`json:"time_start"`
	TimeEnd		time.Time	`json:"time_end"`
}

