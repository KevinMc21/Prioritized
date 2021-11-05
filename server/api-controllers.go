package server

import (
	"Prioritized/v0/scoring"
	"Prioritized/v0/sorting"
	"Prioritized/v0/tasks"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Operation int

const (
	Insert Operation = iota // Inserts a task
	Delete			// Deletes a task
	ReSort			// Provided groupings' tasks gets resorted
	Update			// Edits a task or task grouping
	Anchor			// Sets the fixed attribute for task or task grouping to be true
	Move			// Move a task to a specific time
)

type InsertTaskRequest struct{
	Preference		float64			`json:"time_preference"`		
	TaskGrouping		tasks.TaskGrouping	`json:"task_grouping" validate:"required"`
	InsertTasks 		[]tasks.Task		`json:"insert_tasks" validate:"required"`
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

type SortTaskRequest struct{
	TaskGrouping 		tasks.TaskGrouping	`json:"task_grouping" validate:"required"`
}

func SortTaskHandler(c echo.Context) error {
	body := new (SortTaskRequest)
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