package tasks

import (
	"sort"
	"sync"

	logger "github.com/AbbasRizvi3/GoLangAssignment.git/internal/logging"
)

type TaskQueue struct {
	Tasks []*Task
	Mutex sync.RWMutex
}

func (q *TaskQueue) AddTask(task *Task) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.Tasks = append(q.Tasks, task)
}

func (q *TaskQueue) GetNextTask() *Task {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	sort.Slice(q.Tasks, func(i, j int) bool {
		return q.Tasks[i].Priority > q.Tasks[j].Priority
	})

	for i := 0; i < len(q.Tasks); i++ {
		t := q.Tasks[i]
		t.Mutex.Lock()
		if t.Status == "Pending" {
			t.Status = "InProgress"
			t.Mutex.Unlock()
			logger.Logger.Info().Msgf("Next Task ID: %s fetched for processing", t.ID)
			return t
		}
		t.Mutex.Unlock()
	}

	logger.Logger.Info().Msg("No pending tasks available in the queue")
	return nil
}

// changed
func (q *TaskQueue) LockTask(task *Task) {
	task.Mutex.Lock()
	defer task.Mutex.Unlock()
	task.Status = "Task Locked"
	logger.Logger.Info().Msgf("Task ID: %s locked for processing", task.ID)
}
