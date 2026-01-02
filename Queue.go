package main

import (
	"sort"
	"sync"
)

type TaskQueue struct {
	Tasks []Task
	mutex sync.RWMutex
}

func (q *TaskQueue) AddTask(task Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.Tasks = append(q.Tasks, task)
}

func (q *TaskQueue) GetNextTask() *Task {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	sort.Slice(q.Tasks, func(i, j int) bool {
		return q.Tasks[i].Priority > q.Tasks[j].Priority
	})

	for i := 0; i < len(q.Tasks); i++ {
		if q.Tasks[i].Status == "Pending" {
			return &q.Tasks[i]
		}
	}
	return nil
}

func (q *TaskQueue) LockTask(task *Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	task.Status = "In Progress"
}
