package main

import (
	"fmt"
	"testing"
)

func TestAddTasksToQueue(t *testing.T) {

	Tasks.Tasks = nil

	task1 := Task{
		Name:     "Task1",
		Priority: 1,
		Status:   "Pending",
	}
	task2 := Task{
		Name:     "Task2",
		Priority: 2,
		Status:   "Pending",
	}
	Tasks.AddTask(&task1)
	Tasks.AddTask(&task2)
	if len(Tasks.Tasks) != 2 {
		t.Errorf("Expected 2 tasks in the queue, got %d", len(Tasks.Tasks))
	}
	fmt.Print(Tasks.Tasks)
}

func TestPriorityOrder(t *testing.T) {
	Tasks.Tasks = nil

	task1 := Task{
		Name:     "LowPriorityTask",
		Priority: 1,
		Status:   "Pending",
	}
	task2 := Task{
		Name:     "HighPriorityTask",
		Priority: 5,
		Status:   "Pending",
	}
	Tasks.AddTask(&task1)
	Tasks.AddTask(&task2)

	nextTask := Tasks.GetNextTask()
	if nextTask.Name != "HighPriorityTask" {
		t.Errorf("Expected HighPriorityTask to be returned first, got %s", nextTask.Name)
	}

}
