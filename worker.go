package main

import (
	"context"
	"time"
)

func worker(jobs <-chan *Task, results chan<- *Task) {
	for j := range jobs {
		ctx, err := context.WithTimeout(context.Background(), 5*time.Second)
		if err != nil {
			Logger.Error().Msgf("Error creating context: %v", err)
		}
		j.Process(ctx) // deal with the error returned.
		results <- j
	}
}
