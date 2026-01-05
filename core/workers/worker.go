package workers

import (
	"context"
	"time"

	logger "github.com/AbbasRizvi3/GoLangAssignment.git/pkg/models/loggers"
)

func Worker(tasks *TaskQueue, jobs <-chan struct{}, results chan<- *Task) {
	for range jobs {
		for {
			task := tasks.GetNextTask()
			if task == nil {
				break
			}
			task.Status = "In Progress"
			tasks.LockTask(task)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := task.Process(ctx)
			if err != nil {
				logger.Logger.Error().Msgf("Error processing Task ID: %s, Error: %v", task.ID, err)
			}
			results <- task
		}
	}
}
