package main

import "sort"

type TaskQueue struct {
	Tasks []Task
}

func (q *TaskQueue) AddTask(task Task) {
	q.Tasks = append(q.Tasks, task)
	sort.Slice(q.Tasks, func(i int, j int) bool {
		return q.Tasks[i].Priority > q.Tasks[j].Priority
	})
}

func (q *TaskQueue) GetNextTask() Task {
	if len(q.Tasks) == 0 {
		return Task{}
	}
	nextTask := q.Tasks[0]
	q.Tasks = q.Tasks[1:]
	return nextTask
}
