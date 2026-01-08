package tasks

import (
	"context"
	"time"

	logger "github.com/AbbasRizvi3/GoLangAssignment.git/internal/logging"
)

const (
	workerTimeout = 5 * time.Second
)

func Worker(tasks *TaskQueue, jobs <-chan struct{}, results chan<- *Task) {
	for range jobs {
		for {
			task := tasks.GetNextTask()
			if task == nil {
				logger.Logger.Info().Msg("No pending tasks in the queue, worker is idling")
				break
			}
			tasks.LockTask(task)
			ctx, cancel := context.WithTimeout(context.Background(), workerTimeout)
			err := task.Process(ctx)
			cancel()
			if err != nil {
				logger.Logger.Error().Msgf("Error processing Task ID: %s, Error: %v", task.ID, err)
			}
			results <- task
		}
	}
}
