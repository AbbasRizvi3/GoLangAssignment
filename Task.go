package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Processable interface {
	Process() error
}

type Task struct {
	ID       string
	Name     string
	Priority int
	Status   string // Pending, InProgress, Completed, Failed
	Result   string
}

func (t *Task) Process(ctx context.Context) error {
	rand := rand.Intn(2)
	Logger.Info().Msgf("Processing Task ID: %s, Name: %s", t.ID, t.Name)
	t.Status = "InProgress"
	select {
	case <-ctx.Done():
		{
			t.Status = "Failed"
			t.Result = "Task cancelled / Timeout"
			Logger.Error().Msgf("Task ID: %s cancelled", t.ID)
			return ctx.Err()
		}
	default:
		{
			time.Sleep(3 * time.Second)
			if rand == 1 {
				t.Status = "Completed"
				t.Result = "Task completed successfully"
				Logger.Info().Msgf("Task ID: %s completed successfully", t.ID)
				return nil
			} else {
				t.Status = "Failed"
				t.Result = "Task failed during processing"
				Logger.Error().Msgf("Task ID: %s failed during processing", t.ID)
				return fmt.Errorf("task %s failed during processing", t.Name)
			}
		}
	}

}
