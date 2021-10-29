package sorting

import (
	"Prioritized/v0/tasks"
	"math"
	"sort"
)

func GreedySortWithInsert(t []tasks.Task, insert tasks.Task) []tasks.Task {
	t = append(t, insert)
	return GreedySort(t)
}

func GreedySort(t []tasks.Task) []tasks.Task {
	sortByScore(t)
	
	var sorted []tasks.Task
	if len(t) <= 3 {
		return t
	} else {
		sorted = append(sorted, t[len(t) - 1], t[0])
	}

	mean := meanScore(t)
	min_ptr := 1
	max_ptr := len(t) - 2
	for i := 0; i < len(t) - 2; i++ {
		min_mean := (scoreOf(t[min_ptr]) + scoreOf(sorted[len(sorted) - 1]) + scoreOf(sorted[len(sorted) - 2]))/3
		max_mean := (scoreOf(t[max_ptr]) + scoreOf(sorted[len(sorted) - 1]) + scoreOf(sorted[len(sorted) - 2]))/3

		if math.Abs(min_mean - mean) <= math.Abs(max_mean - mean) {
			sorted = append(sorted, t[min_ptr])
			min_ptr++
		} else {
			sorted = append(sorted, t[max_ptr])
			max_ptr--
		}
	}

	return sorted
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

type SortBy []tasks.Task

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].CurrentScore < a[j].CurrentScore }

