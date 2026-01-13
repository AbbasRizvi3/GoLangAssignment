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

func fetchTask(t *Task) (*Task, bool) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	if t.Status == "Pending" {
		t.Status = "InProgress"
		return t, true
	}
	return nil, false
}

func (q *TaskQueue) GetNextTask() *Task {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	sort.Slice(q.Tasks, func(i, j int) bool {
		return q.Tasks[i].Priority > q.Tasks[j].Priority
	})

	for i := 0; i < len(q.Tasks); i++ {
		if t, ok := fetchTask(q.Tasks[i]); ok {
			logger.Logger.Info().Msgf("Next Task ID: %s fetched for processing", t.ID)
			return t
		}
	}

	logger.Logger.Info().Msg("No pending tasks available in the queue")
	return nil
}
