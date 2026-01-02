package main

import (
	"context"
	"time"
)

func worker(jobs <-chan struct{}, results chan<- *Task) {
	for range jobs {
		for {
			task := Tasks.GetNextTask()
			if task == nil {
				break
			}
			task.Status = "In Progress"
			Tasks.LockTask(task)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := task.Process(ctx)
			if err != nil {
				Logger.Error().Msgf("Error processing Task ID: %s, Error: %v", task.ID, err)
			}
			results <- task
		}
	}
}
