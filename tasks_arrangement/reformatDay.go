package tasksarrangement

import (
	"Prioritized/v0/GeneticAlgo"
	"Prioritized/v0/tasks"
	"fmt"
	"time"
)

func ReformatDay(task []GeneticAlgo.Day) []tasks.Task {
	var sortedTask []tasks.Task

	givenTime := [][]int{{9, 0}, {9, 30}, {10, 00}, {10, 30}, {11, 00}, {11, 30}, {13, 00}, {13, 30}}

	currentDatePointer := 0
	for _, i := range task {
		var tempTask tasks.Task
		sortArr := sortTask(i.Items)
		checkcursor := time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() + currentDatePointer + 1), 0, 0, 0, 0, time.Local).Weekday().String()
		for checkcursor == "Saturday" || checkcursor == "Sunday" {
			currentDatePointer++
		}

		for h, t := range sortArr {
			dateNow := time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() + currentDatePointer + 1), givenTime[h][0], givenTime[h][1], 0, 0, time.Local)
			if t.EstimatedTime.String() == "0s" {
				break
			}
			tempTask.Name = t.Name
			fmt.Println("Print Name : ", tempTask.Name, t.Name)
			tempTask.CurrentScore = t.CurrentScore
			tempTask.Timeline = t.Timeline
			tempTask.AssignedTime.TimeStart = dateNow
			timeAdd, _ := time.ParseDuration("30m")
			tempTask.AssignedTime.TimeEnd = dateNow.Add(timeAdd)
			tempTask.EstimatedTime = t.EstimatedTime
			sortedTask = append(sortedTask, tempTask)
		}
		currentDatePointer++
	}

	for _, i := range sortedTask {
		fmt.Println("CHECKING! : ", i.AssignedTime.TimeStart.String(), i.AssignedTime.TimeEnd.String())
	}
	return sortedTask
}

//yyyy-mm-ddThh:mm:ss+07:00

func sortTask(taskArr [8]tasks.Task) [8]tasks.Task {

	var sortedTask [8]tasks.Task
	checkMap := make(map[string]bool)
	pos := 0

	for _, i := range taskArr {
		if !checkMap[i.Name] {
			for _, j := range taskArr {
				if i.Name == j.Name {
					sortedTask[pos] = j
					pos++
				}
			}
			checkMap[i.Name] = true
		}
	}

	return sortedTask
}
