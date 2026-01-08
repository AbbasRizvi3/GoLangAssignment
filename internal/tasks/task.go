package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	logger "github.com/AbbasRizvi3/GoLangAssignment.git/internal/logging"
)

type Processable interface {
	Process(ctx context.Context) error
}

type Task struct {
	ID       string
	Name     string
	Priority int
	Status   string // Pending, InProgress, Completed, Failed
	Result   string
	Mutex    sync.RWMutex `json:"-"`
}

func (t *Task) Process(ctx context.Context) error {
	t.Mutex.Lock()
	t.Status = "InProgress"
	t.Mutex.Unlock()

	rand := rand.Intn(2)
	logger.Logger.Info().Msgf("Processing Task ID: %s, Name: %s", t.ID, t.Name)
	select {
	case <-ctx.Done():
		{
			t.Mutex.Lock()
			t.Status = "Failed"
			t.Result = "Task cancelled / Timeout"
			t.Mutex.Unlock()
			logger.Logger.Error().Msgf("Task ID: %s cancelled", t.ID)
			return ctx.Err()
		}
	default:
		{
			time.Sleep(3 * time.Second)
			if rand == 1 {
				t.Mutex.Lock()
				t.Status = "Completed"
				t.Result = "Task completed successfully"
				t.Mutex.Unlock()
				logger.Logger.Info().Msgf("Task ID: %s completed successfully", t.ID)
				return nil
			} else {
				t.Mutex.Lock()
				t.Status = "Failed"
				t.Result = "Task failed during processing"
				t.Mutex.Unlock()
				logger.Logger.Error().Msgf("Task ID: %s failed during processing", t.ID)
				return fmt.Errorf("task %s failed during processing", t.Name)
			}
		}
	}

}
