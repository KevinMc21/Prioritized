package sorting

import (
	"Prioritized/v0/tasks"
	"fmt"
	"math"
	"sort"
	"time"
)

func GreedySortWithInsert(t tasks.TaskGrouping, insert []tasks.Task) []tasks.Task {
	// for _, task := range insert {
	// 	task.AssignedTime.TimeStart = time.Time{}
	// 	task.AssignedTime.TimeEnd = time.Time{}
	// }

	t.Tasks = append(t.Tasks, insert...)
	return GreedySort(t)
}

func GreedySort(t tasks.TaskGrouping) []tasks.Task {
	sortByScore(t.Tasks)
	
	var sorted []tasks.Task
	if len(t.Tasks) <= 3 {
		return t.Tasks
	} else {
		sorted = append(sorted, t.Tasks[len(t.Tasks) - 1], t.Tasks[0])
	}

	mean := meanScore(t.Tasks)
	min_ptr := 1
	max_ptr := len(t.Tasks) - 2
	for i := 0; i < len(t.Tasks) - 2; i++ {
		min_mean := (scoreOf(t.Tasks[min_ptr]) + scoreOf(sorted[len(sorted) - 1]) + scoreOf(sorted[len(sorted) - 2]))/3
		max_mean := (scoreOf(t.Tasks[max_ptr]) + scoreOf(sorted[len(sorted) - 1]) + scoreOf(sorted[len(sorted) - 2]))/3

		if math.Abs(min_mean - mean) <= math.Abs(max_mean - mean) {
			sorted = append(sorted, t.Tasks[min_ptr])
			min_ptr++
		} else {
			sorted = append(sorted, t.Tasks[max_ptr])
			max_ptr--
		}
	}

	assigned := AssignTimes(sorted, t.TimeRanges, t.Weekdays)

	return assigned
}

func sortByScore(t []tasks.Task) {
	sort.Sort(SortBy(t))
}

func meanScore(t []tasks.Task) float64 {
	sum := 0.0
	for _, task := range t {
		sum += task.CurrentScore
	}

	return sum/float64(len(t))
}

func scoreOf(t tasks.Task) float64 {
	return t.CurrentScore
}

// AssignTimes will assign times to tasks inside of t, assigned times will be between the time periods specified by periods.
// Tasks will not be assigned on days not included in weekdaysAvailable. To include all days, pass in []time.Weekday{}
func AssignTimes(t []tasks.Task, periods []tasks.Period, weekdaysAvailable []time.Weekday) (assigned []tasks.Task) {
	timeToInsert := time.Now().In(periods[0].TimeStart.Location())	
	pastDeadline := []int{}

	for i, task := range t {
		timeOnTaskDuration, _ := time.ParseDuration("30m")
		timeToInsert = nextTimeAfter(timeToInsert, periods, task.EstimatedTime.Round(timeOnTaskDuration))
		timeToInsert = moveDay(timeToInsert, weekdaysAvailable)

		if !withinTimeline(timeToInsert, task.Timeline) {
			pastDeadline = append(pastDeadline, i)
		}

		task.AssignedTime.TimeStart = timeToInsert
		timeToInsert = timeToInsert.Add(task.EstimatedTime)
		task.AssignedTime.TimeEnd = timeToInsert
		assigned = append(assigned, task)
	}

	previouSwap := -1
	found := false
	for _, task := range pastDeadline {
		difference := math.Inf(1)
		swapIndex := 0
		for index, swapTask := range t {
			if !withinTimeline(swapTask.AssignedTime.TimeStart,t[task].Timeline) {
				continue
			}

			currentDifference := math.Abs(t[task].CurrentScore - swapTask.CurrentScore)
			if difference > currentDifference && index != previouSwap {
				difference = currentDifference
				swapIndex = index
				found = true
			}
		}
		if !found {
			if previouSwap - 1 >= 0 {
				swapIndex = previouSwap - 1
			} else if previouSwap + 1 < len(t) {
				swapIndex = previouSwap + 1
			}
		}

		assigned[task].AssignedTime, assigned[swapIndex].AssignedTime = assigned[swapIndex].AssignedTime, assigned[task].AssignedTime

		assigned[task], assigned[swapIndex] = assigned[swapIndex], assigned[task]
		previouSwap = swapIndex
	}

	return
}

func moveDay(timeToChange time.Time, weekdays []time.Weekday) time.Time {
	if len(weekdays) == 0 {
		return timeToChange
	}

	currentWeekday := timeToChange.Weekday()

	weekdaysAvailable := make(map[time.Weekday]bool)
	for _, weekday := range weekdays {
		if weekday == currentWeekday {
			return timeToChange
		}
		weekdaysAvailable[weekday] = true
	}

	moveDays := currentWeekday
	for i := 0; i < 7; i++ {
		moveDays++
		if _, ok := weekdaysAvailable[moveDays % 7]; ok {
			duration, err := time.ParseDuration(fmt.Sprintf("%dh", (moveDays - currentWeekday) * 24))
			if err != nil {
				return timeToChange
			}

			return timeToChange.Add(duration)
		}
	}


	return timeToChange
}

// Returns the beginning of the next valid time period
func nextTimeAfter(t time.Time, periods []tasks.Period, durationOffset time.Duration) time.Time {
	if timeBetween(t, periods) && timeBetween(t.Add(durationOffset), periods) {
		return t
	}

	h1, _, _ := t.Clock()

	for _, period := range periods {
		hourStart, _, _ := period.TimeStart.Clock()

		if hourStart > h1 {
			return time.Date(t.Year(), t.Month(), t.Day(), period.TimeStart.Hour(), period.TimeStart.Minute(), period.TimeStart.Second(), 0, t.Location())
		}
	}

	return time.Date(t.Year(), t.Month(), t.Day() + 1, periods[0].TimeStart.Hour(), periods[0].TimeStart.Minute(), periods[0].TimeStart.Second(), 0, t.Location())
}

func withinTimeline(t time.Time, timeline tasks.Period) (bool) {
	if t.After(timeline.TimeStart) {
		if timeline.TimeEnd.IsZero() {
			return true
		} else if t.Before(timeline.TimeEnd) {
			return true
		}
	}

	return false
}

func timeBetween(t time.Time, periods []tasks.Period) (bool) {
	for _, period := range periods {
		period.TimeStart = period.TimeStart.In(t.Location())
		period.TimeEnd = period.TimeEnd.In(t.Location())

		pStartHour, pStartMinute, _ := period.TimeStart.Clock()
		pEndHour, pEndMinute, _ := period.TimeEnd.Clock()
		
		tHour, tMinute, _ := t.Clock()

		if (pStartHour == tHour && tMinute < pStartMinute) || (pEndHour == tHour && tMinute > pEndMinute) {
			return false
		} else if pStartHour < tHour && pEndHour > tHour {
			return true
		}
	}

	return false
}

type SortBy []tasks.Task

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].CurrentScore < a[j].CurrentScore }

