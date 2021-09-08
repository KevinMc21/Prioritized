package tasks

import (
	"time"
)

type Time string
type Datetime string

type DatetimePeriod struct{
	TimeStart 	Datetime	`json:"time_start"`
	TimeEnd		Datetime	`json:"time_end"`
}

type TimePeriod struct{
	TimeStart	Time
	TimeEnd		Time
}

func (dt Datetime) Parse() (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, string(dt))
	if err != nil {
		t = time.Time{}		
	}

	return t.UTC(), err
}


func (t Time) Parse() (time.Time, error) {
	parsed, err := time.Parse("15:04:05Z07:00", string(t))
	if err != nil {
		return time.Time{}, err
	}

	return parsed.UTC(), nil
}





