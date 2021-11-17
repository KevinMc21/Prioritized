package GeneticAlgo

import (
	"Prioritized/v0/scoring"
	"Prioritized/v0/tasks"
	"fmt"
	"math/rand"
	"time"
)

type Day struct {
	Items       [8]tasks.Task
	ItemsMap    map[string]time.Duration
	WeightCoef  float64
	Fitness     float64
	TotatEnergy float64
}

func NewBag(taskArr []tasks.Task) *Day {
	var tempArr []tasks.Task
	tempArr = append(tempArr, taskArr...)
	B := new(Day)

	B.ItemsMap = make(map[string]time.Duration)
	for _, i := range tempArr {
		B.ItemsMap[i.Name] = i.EstimatedTime
		fmt.Println("Estimate time : ", B.ItemsMap[i.Name])
	}

	for i := 0; i < 8; i++ {
		if len(tempArr) > 0 {
			choosedTaskIndex := 0
			rand.Seed(time.Now().UnixNano())

			if len(tempArr) > 1 {
				choosedTaskIndex = rand.Intn(len(tempArr) - 1)
			} else if len(tempArr) == 1 {
				choosedTaskIndex = 0
			}

			B.Items[i] = tempArr[choosedTaskIndex]
			fmt.Println("End Gene", len(tempArr))
			tempArr = deductedHour(tempArr, 30, tempArr[choosedTaskIndex].Name)
		} else {
			B.Items[i] = tasks.Task{}
		}

	}

	// fmt.Println("End Bag Gen")
	return &Day{Items: B.Items, ItemsMap: B.ItemsMap}
}

func (D *Day) CalFitness() {
	D.TotatEnergy = 3000
	D.Fitness = 0
	if D.CheckSlot() {
		D.Fitness = 0
	} else {
		for m, i := range D.Items {
			D.Fitness += i.CurrentScore
			D.TotatEnergy = D.TotatEnergy - (i.CurrentScore * float64(1/float64(8-m)))
			if D.TotatEnergy <= 0 {
				D.Fitness = 0
				break
			}
		}
	}
}

func (D *Day) CheckSlot() bool {
	checkMap := make(map[string]int)

	for _, i := range D.Items {
		checkMap[i.Name]++
		var timeD time.Duration
		timeM, _ := time.ParseDuration("30m")

		for j := 1; j <= checkMap[i.Name]; j++ {
			timeD += timeM
		}

		if timeD > (D.ItemsMap[i.Name]) {
			for _, i := range D.Items {
				fmt.Print(i.Name)
			}
			fmt.Println(" Time : ", timeD, D.ItemsMap[i.Name], i.Name)
			return true
		}
	}
	return false
}

func deductedHour(DTask []tasks.Task, preferedTime int, Name string) []tasks.Task {
	choosenIndex := tasks.SearchTask(Name, &DTask)
	pointerTask := DTask[choosenIndex]
	thistime, _ := time.ParseDuration("30m")
	pointerTask.EstimatedTime = pointerTask.EstimatedTime - thistime

	pointerTask.CurrentScore = scoring.GiveScore(pointerTask.EstimatedTime, 30, pointerTask.WeightCoef, pointerTask.WeightCoef)

	DTask[choosenIndex].EstimatedTime = pointerTask.EstimatedTime
	DTask[choosenIndex].CurrentScore = pointerTask.CurrentScore

	if lessThan(DTask[choosenIndex].EstimatedTime) {
		DTask = remove(DTask, choosenIndex)
	}

	return DTask
}

func lessThan(t time.Duration) bool {
	return t.Hours() <= 0 && t.Seconds() <= 0 && t.Minutes() <= 0
}

func remove(s []tasks.Task, i int) []tasks.Task {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
