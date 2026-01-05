package workers

import (
	"sort"
	"sync"
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
	q.Mutex.RLock()
	defer q.Mutex.RUnlock()
	sort.Slice(q.Tasks, func(i, j int) bool {
		return q.Tasks[i].Priority > q.Tasks[j].Priority
	})

	for i := 0; i < len(q.Tasks); i++ {
		if q.Tasks[i].Status == "Pending" {
			return q.Tasks[i]
		}
	}
	return nil
}

func (q *TaskQueue) LockTask(task *Task) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	task.Status = "In Progress"
}
