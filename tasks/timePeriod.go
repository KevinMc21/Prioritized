package tasks

import "time"

type Period struct {
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
}

func ValidWeekday(wd time.Weekday) bool {
	switch wd {
	case time.Sunday:
		return true
	case time.Monday:
		return true
	case time.Tuesday:
		return true
	case time.Wednesday:
		return true
	case time.Thursday:
		return true
	case time.Friday:
		return true
	case time.Saturday:
		return true
	default:
		return false
	}
}

type SortByWeekday []time.Weekday

func (a SortByWeekday) Len() int           { return len(a) }
func (a SortByWeekday) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByWeekday) Less(i, j int) bool { return int(a[i]) < int(a[j]) }
