package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/core/app"
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
)

const (
	totalTasks  = 20
	workerCount = 4
	timeoutSec  = 6
)

func TestWorkersProcessAllTasksUsingAppWorkers(t *testing.T) {
	app.Tasks.Tasks = nil

	for w := 0; w < workerCount; w++ {
		go tasks.Worker(&app.Tasks, app.TaskChannel, app.ResultChannel)
	}

	time.Sleep(10 * time.Millisecond)

	for i := 0; i < totalTasks; i++ {
		task := tasks.Task{
			Name:     fmt.Sprintf("app-wp-task-%02d", i),
			Priority: i % 3,
			ID:       fmt.Sprintf("task-%02d", i),
			Status:   "Pending",
		}
		app.Tasks.AddTask(&task)

		app.TaskChannel <- struct{}{}
	}

	received := make(map[string]bool)
	deadline := time.After(time.Duration(timeoutSec) * time.Second)

	for len(received) < totalTasks {
		select {
		case res := <-app.ResultChannel:
			received[res.Name] = true
		case <-deadline:
			t.Fatalf("timeout waiting for results: got %d/%d", len(received), totalTasks)
		}
	}

	if len(received) != totalTasks {
		t.Fatalf("expected %d processed tasks, got %d", totalTasks, len(received))
	}
}

func TestWorkersAvoidDuplicateProcessingUsingAppWorkers(t *testing.T) {
	app.Tasks.Tasks = nil

	for w := 0; w < workerCount; w++ {
		go tasks.Worker(&app.Tasks, app.TaskChannel, app.ResultChannel)
	}
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < totalTasks; i++ {
		task := tasks.Task{
			Name:     fmt.Sprintf("dup-task-%02d", i),
			Priority: i % 2,
			ID:       fmt.Sprintf("dup-%02d", i),
			Status:   "Pending",
		}
		app.Tasks.AddTask(&task)
		app.TaskChannel <- struct{}{}
	}

	var mu sync.Mutex
	counts := make(map[string]int)
	deadline := time.After(time.Duration(timeoutSec) * time.Second)
	processed := 0

	for processed < totalTasks {
		select {
		case res := <-app.ResultChannel:
			mu.Lock()
			counts[res.Name]++
			if counts[res.Name] > 1 {
				mu.Unlock()
				t.Fatalf("task processed more than once: %s", res.Name)
			}
			mu.Unlock()
			processed++
		case <-deadline:
			t.Fatalf("timeout waiting for results: processed %d/%d", processed, totalTasks)
		}
	}

	if len(counts) != totalTasks {
		t.Fatalf("expected %d unique processed tasks, got %d", totalTasks, len(counts))
	}
}
