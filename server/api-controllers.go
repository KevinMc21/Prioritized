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
	Preferences		float64			`json:"preference"`		
	TaskGrouping		tasks.TaskGrouping	`json:"task_grouping" validate:"required"`
	Task 			tasks.Task		`json:"task" validate:"required"`
}

func InsertTaskHandler(c echo.Context) error {
	body := new(InsertTaskRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task := body.Task
	score := scoring.GiveScore(task.EstimatedTime, body.Preferences, body.Task.WeightCoef, body.TaskGrouping.WeightCoef)
	task.CurrentScore = score

	if len(body.TaskGrouping.Tasks) == 0 {
		inserted := append(body.TaskGrouping.Tasks, task)
		return c.JSON(http.StatusOK, inserted)
	}

	sorted := sorting.GreedySortWithInsert(body.TaskGrouping.Tasks, task)

	return c.JSON(http.StatusOK, sorted)
}