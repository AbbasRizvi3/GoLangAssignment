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
	totalTasks  = 5
	workerCount = 4
	timeoutSec  = 6
)

var ActiveWorkers = 0
var mutex sync.Mutex

func TestWorkersProcessAllTasksUsingAppWorkers(t *testing.T) {
	app.Tasks.Tasks = nil

	go func() {
		for range app.TaskChannel {
			for {
				mutex.Lock()
				if ActiveWorkers < workerCount {
					ActiveWorkers++
					mutex.Unlock()
					go tasks.Worker(&app.Tasks, &app.ResultSlice, &ActiveWorkers, &mutex)
					break
				}
				mutex.Unlock()
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	for i := 0; i < totalTasks; i++ {
		task := &tasks.Task{
			Name:     fmt.Sprintf("app-wp-task-%02d", i),
			Priority: i % 3,
			ID:       fmt.Sprintf("task-%02d", i),
			Status:   "Pending",
		}
		app.Tasks.AddTask(task)

		app.TaskChannel <- struct{}{}
	}

	time.Sleep(10 * time.Second)

	if len(app.ResultSlice) != totalTasks {
		t.Fatalf("expected %d processed tasks, got %d", totalTasks, len(app.ResultSlice))
	}
}
func TestWorkersAvoidDuplicateProcessingUsingAppWorkers(t *testing.T) {

	app.Tasks.Tasks = nil
	app.ResultSlice = nil
	ActiveWorkers = 0

	go func() {
		for range app.TaskChannel {
			for {
				mutex.Lock()
				if ActiveWorkers < workerCount {
					ActiveWorkers++
					mutex.Unlock()
					go tasks.Worker(&app.Tasks, &app.ResultSlice, &ActiveWorkers, &mutex)
					break
				}
				mutex.Unlock()
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()

	for i := 0; i < totalTasks; i++ {
		task := &tasks.Task{
			Name:     fmt.Sprintf("dup-task-%02d", i),
			Priority: i % 2,
			ID:       fmt.Sprintf("dup-%02d", i),
			Status:   "Pending",
		}
		app.Tasks.AddTask(task)
		app.TaskChannel <- struct{}{}
	}
	time.Sleep(10 * time.Second)

	seen := make(map[string]bool)

	if len(app.ResultSlice) != totalTasks {
		t.Fatalf("Expected %d results, but got %d", totalTasks, len(app.ResultSlice))
	}

	for i := range app.ResultSlice {
		res := app.ResultSlice[i]
		if seen[res.ID] {
			t.Fatalf("FAILED: Task %s was processed more than once!", res.ID)
		}
		seen[res.ID] = true
	}

	t.Logf("Success: All %d tasks processed uniquely.", len(seen))
}
