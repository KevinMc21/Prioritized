package server

import (
	"Prioritized/v0/GeneticAlgo"
	"Prioritized/v0/scoring"
	"Prioritized/v0/sorting"
	"Prioritized/v0/tasks"
	tasksarrangement "Prioritized/v0/tasks_arrangement"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Operation int

const (
	Insert Operation = iota // Inserts a task
	Delete                  // Deletes a task
	ReSort                  // Provided groupings' tasks gets resorted
	Update                  // Edits a task or task grouping
	Anchor                  // Sets the fixed attribute for task or task grouping to be true
	Move                    // Move a task to a specific time
)

type InsertTaskRequest struct {
	Preference   float64            `json:"time_preference"`
	TaskGrouping tasks.TaskGrouping `json:"task_grouping" validate:"required"`
	InsertTasks  []tasks.Task       `json:"insert_tasks" validate:"required"`
}

func InsertTaskHandler(c echo.Context) error {
	body := new(InsertTaskRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	scoredTasks := []tasks.Task{}
	for _, task := range body.InsertTasks {
		score := scoring.GiveScore(task.EstimatedTime, body.Preference, task.WeightCoef, body.TaskGrouping.WeightCoef)
		task.CurrentScore = score
		scoredTasks = append(scoredTasks, task)
	}

	if len(body.TaskGrouping.Tasks) == 0 {
		inserted := append(body.TaskGrouping.Tasks, scoredTasks...)
		return c.JSON(http.StatusOK, inserted)
	}

	sorted := sorting.GreedySortWithInsert(body.TaskGrouping, scoredTasks)

	return c.JSON(http.StatusOK, sorted)
}

type SortTaskRequest struct {
	TaskGrouping tasks.TaskGrouping `json:"task_grouping" validate:"required"`
}

func SortTaskHandler(c echo.Context) error {
	body := new(SortTaskRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(body.TaskGrouping.Tasks) == 0 {
		return c.JSON(http.StatusOK, body.TaskGrouping)
	}

	sorted := sorting.GreedySort(body.TaskGrouping)

	return c.JSON(http.StatusOK, sorted)
}

func InsertTaskGeneticHandler(c echo.Context) error {
	body := new(SortTaskRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var allTask []tasks.Task = []tasks.Task{}

	allTask = append(allTask, tasksarrangement.ScoreTask(body.TaskGrouping, time.Now())...)

	for _, i := range allTask {
		i.EstimatedTime, _ = time.ParseDuration("1h30m0s")
		fmt.Println(i.Name)
	}

	var TaskAsignment []GeneticAlgo.Day
	var tempTask []tasks.Task
	tempTask = append(tempTask, allTask...)

	for len(tempTask) != 0 {
		Ouput, leftover := GeneticAlgo.RunGeneticAlgorithm(tempTask)

		TaskAsignment = append(TaskAsignment, Ouput)

		tempTask = []tasks.Task{}
		tempTask = append(tempTask, leftover...)

	}

	formatted := tasksarrangement.ReformatDay(TaskAsignment)

	return c.JSON(http.StatusOK, formatted)
}
