package tasks

import "time"

type TaskCategory struct{
	Name		string
	Difficulty	float32
	TimeRanges	[]time.Time
}