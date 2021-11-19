package tasksarrangement

import (
	sr "Prioritized/v0/scoring"
	"Prioritized/v0/tasks"
	"fmt"
	"time"
)

func ScoreTask(t tasks.TaskGrouping, time time.Time) []tasks.Task {

	curtime := time
	var scoredTask []tasks.Task

	for _, ts := range t.Tasks {
		dateDiff := TimeDiff(curtime, ts)
		ts.CurrentScore = sr.GiveScore(ts.EstimatedTime, 30, 1, t.WeightCoef) * (1 + dateDiff)
		fmt.Println("AFTER SCORING : ", ts.CurrentScore)
		fmt.Println("Pure Score : ", sr.GiveScore(ts.EstimatedTime, 30, 1, t.WeightCoef))
		scoredTask = append(scoredTask, ts)
	}

	return scoredTask
}

func TimeDiff(curtime time.Time, t tasks.Task) float64 {
	deadline := t.Timeline.TimeEnd
	yearC, monthC, dayC := curtime.Date()
	yearA, monthA, dayA := deadline.Date()

	if !t.AssignedTime.TimeStart.Equal(curtime) {
		var diff float64 = 2 / (float64(yearA-yearC) + float64(monthA-monthC) + float64(dayA-dayC))
		return diff
	} else {
		return 0.0
	}
}
