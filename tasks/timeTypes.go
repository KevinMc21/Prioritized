package tasks

import "time"

type Time struct{
	time		time.Time
}

func (t Time) Get() (hour int, min int, sec int) {
	return t.time.Clock()
}
