package tasks

import (
	"context"
	"sync"
	"time"

	logger "github.com/AbbasRizvi3/GoLangAssignment.git/internal/logging"
)

const (
	workerTimeout = 5 * time.Second
)

func decrementActiveWorkers(ActiveWorkers *int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	*ActiveWorkers--
}

func addResult(task *Task, tasks *TaskQueue, results *[]Task) {
	task.Mutex.Lock()
	defer task.Mutex.Unlock()
	index := -1
	for i, t := range tasks.Tasks {
		if t.ID == task.ID {
			index = i
			break
		}
	}
	if index != -1 {
		tasks.Tasks = append(tasks.Tasks[:index], tasks.Tasks[index+1:]...)
		*results = append(*results, *task)
	}

}

func Worker(tasks *TaskQueue, results *[]Task, ActiveWorkers *int, mutex *sync.Mutex) {

	task := tasks.GetNextTask()
	if task == nil {
		logger.Logger.Info().Msg("No pending tasks in the queue, worker is idling")
	}

	ctx, cancel := context.WithTimeout(context.Background(), workerTimeout)
	err := task.Process(ctx)
	addResult(task, tasks, results)
	if err != nil {
		logger.Logger.Error().Msgf("Error processing Task ID: %s, Error: %v", task.ID, err)
	}

	decrementActiveWorkers(ActiveWorkers, mutex)
	logger.Logger.Info().Msgf("Worker finished processing Task ID: %s, Active Workers: %d", task.ID, *ActiveWorkers)
	defer cancel()

}
