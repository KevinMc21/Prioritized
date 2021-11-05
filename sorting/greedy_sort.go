package sorting

import (
	"Prioritized/v0/loggers/debug"
	"Prioritized/v0/tasks"
	"math"
	"sort"
	"time"
)

func GreedySortWithInsert(t tasks.TaskGrouping, insert tasks.Task) []tasks.Task {
	t.Tasks = append(t.Tasks, insert)
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

	assigned := assignTimes(sorted, t.TimeRanges)

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

func assignTimes(t []tasks.Task, periods []tasks.Period) (assigned []tasks.Task) {
	timeToInsert := time.Now().In(periods[0].TimeStart.Location())	
	pastDeadline := []int{}

	for i, task := range t {
		timeToInsert = nextTimeAfter(timeToInsert, periods)

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

		temp := t[task].AssignedTime
		t[task].AssignedTime = t[swapIndex].AssignedTime
		t[swapIndex].AssignedTime = temp
		
		assigned[task] = t[swapIndex]
		assigned[swapIndex] = t[task]
		debug.GetDebugLogger().Println("swapped tasks", t[task], t[swapIndex])
		previouSwap = swapIndex
	}

	return
}

// Returns the beginning of the next valid time period
func nextTimeAfter(t time.Time, periods []tasks.Period) time.Time {
	if timeBetween(t, periods) {
		return t
	}

	h1, _, _ := t.Clock()

	for _, period := range periods {
		hourStart, _, _ := period.TimeStart.Clock()

		if hourStart > h1 {
			return time.Date(t.Year(), t.Month(), t.Day() + 1, period.TimeStart.Hour(), period.TimeStart.Minute(), period.TimeStart.Second(), 0, t.Location())
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

