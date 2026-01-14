package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	logger "github.com/AbbasRizvi3/GoLangAssignment.git/internal/logging"
)

const (
	randLimit = 10
	sleepTime = 3 * time.Second
)

type Processable interface {
	Process(ctx context.Context) error
}

type Task struct {
	ID       string
	Name     string
	Priority int
	Status   string
	Result   string
	Mutex    sync.RWMutex `json:"-"`
}

func changeStatusToInProgress(t *Task) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Status = "InProgress"
}

func caseTimeout(t *Task) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Status = "Failed"
	t.Result = "Timeout occurred during processing"
}

func caseTaskCompleted(t *Task) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Status = "Completed"
	t.Result = "Task completed successfully"
}

func caseTaskFailed(t *Task) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Status = "Failed"
	t.Result = "Task failed during processing"
}

func (t *Task) Process(ctx context.Context) error {
	changeStatusToInProgress(t)

	rand := rand.Intn(randLimit)
	logger.Logger.Info().Msgf("Processing Task ID: %s, Name: %s", t.ID, t.Name)

	if rand == 0 {
		caseTaskFailed(t)
		logger.Logger.Error().Msgf("Task ID: %s failed during processing", t.ID)
		return fmt.Errorf("task %s failed during processing", t.Name)
	}

	select {
	case <-ctx.Done():
		{
			caseTimeout(t)
			logger.Logger.Error().Msgf("Task ID: %s cancelled", t.ID)
			return ctx.Err()
		}
	case <-time.After(time.Duration(rand) * time.Second):
		caseTaskCompleted(t)
		logger.Logger.Info().Msgf("Task ID: %s completed successfully", t.ID)
		return nil
	}

}
