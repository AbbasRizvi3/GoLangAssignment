package tests

import (
	"context"
	"testing"
	"time"

	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
)

func TestTaskProcessingWithoutIssues(t *testing.T) {
	testTask := &tasks.Task{
		Name:     "Test Task",
		Priority: 1,
		Status:   "Pending",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := testTask.Process(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if testTask.Status != "Completed" && testTask.Status != "Failed" {
		t.Errorf("Expected status to be Completed or Failed, got %s", testTask.Status)
	}
}

func TestTaskProcessingWithTimeout(t *testing.T) {
	testTask := &tasks.Task{
		Name:     "Test Task with Timeout",
		Priority: 1,
		Status:   "Pending",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	time.Sleep(2 * time.Nanosecond)
	err := testTask.Process(ctx)
	if err == nil {
		t.Errorf("Expected timeout error, got nil")
	}
	if testTask.Status != "Failed" {
		t.Errorf("Expected status to be Failed due to timeout, got %s", testTask.Status)
	}
}

func TestTaskProcessRandomness(t *testing.T) {
	completed, failed := 0, 0

	for i := 0; i < 20; i++ {
		task := &tasks.Task{
			Name:     "Random Task",
			Priority: 1,
			Status:   "Pending",
		}

		err := task.Process(context.Background())

		if task.Status == "Completed" {
			completed++
		} else if task.Status == "Failed" {
			failed++
		} else {
			t.Errorf("Unexpected status: %s", task.Status)
		}

		if task.Status == "Completed" && err != nil {
			t.Errorf("Expected nil error for Completed task")
		}
		if task.Status == "Failed" && err == nil {
			t.Errorf("Expected error for Failed task")
		}
	}
	if completed == 0 || failed == 0 {
		t.Errorf("Randomness not observed: Completed=%d, Failed=%d", completed, failed)
	}
}
