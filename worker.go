package main

import (
	"context"
	"time"
)

func worker(jobs <-chan *Task, results chan<- *Task) {
	for j := range jobs {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		j.Process(ctx) // deal with the error returned.
		results <- j
	}
}
