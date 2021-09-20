package tasks

import (
	"time"
)

// Type to hold time only data
type Time string

// Type to hold Datetimes
type Datetime string

// TimeStart specifies when the period starts and TimeEnd specifies when the period ends. Stores Datetime types
type DatetimePeriod struct{
	TimeStart 	Datetime	`json:"datetime_start"`
	TimeEnd		Datetime	`json:"datetime_end"`
}

// TimeStart specifies when the period starts and TimeEnd specifies when the period ends. Stores Time types
type TimePeriod struct{
	TimeStart	Time		`json:"time_start"`
	TimeEnd		Time		`json:"time_end"`
}

// Parse Datetime string and return time.Time object
func (dt Datetime) Parse() (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, string(dt))
	if err != nil {
		t = time.Time{}		
	}

	return t.UTC(), err
}

// Parse Time string and return time.Time object
func (t Time) Parse() (time.Time, error) {
	parsed, err := time.Parse("15:04:05Z07:00", string(t))
	if err != nil {
		return time.Time{}, err
	}

	return parsed.UTC(), nil
}





