package tests

import (
	"fmt"
	"testing"

	"github.com/AbbasRizvi3/GoLangAssignment.git/core/app"
	"github.com/AbbasRizvi3/GoLangAssignment.git/core/workers"
)

func TestAddTasksToQueue(t *testing.T) {

	app.Tasks.Tasks = nil

	task1 := workers.Task{
		Name:     "Task1",
		Priority: 1,
		Status:   "Pending",
	}
	task2 := workers.Task{
		Name:     "Task2",
		Priority: 2,
		Status:   "Pending",
	}
	app.Tasks.AddTask(&task1)
	app.Tasks.AddTask(&task2)
	if len(app.Tasks.Tasks) != 2 {
		t.Errorf("Expected 2 tasks in the queue, got %d", len(app.Tasks.Tasks))
	}
	fmt.Print(app.Tasks.Tasks)
}

func TestPriorityOrder(t *testing.T) {
	app.Tasks.Tasks = nil
	task1 := workers.Task{
		Name:     "LowPriorityTask",
		Priority: 1,
		Status:   "Pending",
	}
	task2 := workers.Task{
		Name:     "HighPriorityTask",
		Priority: 5,
		Status:   "Pending",
	}
	app.Tasks.AddTask(&task1)
	app.Tasks.AddTask(&task2)

	nextTask := app.Tasks.GetNextTask()
	if nextTask.Name != "HighPriorityTask" {
		t.Errorf("Expected HighPriorityTask to be returned first, got %s", nextTask.Name)
	}

}
