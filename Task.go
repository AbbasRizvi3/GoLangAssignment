package main

import (
	"errors"
	"math/rand"
	"time"
)

type Processable interface {
	Process() error
}

type Task struct {
	ID       string
	Name     string
	Priority int
	Status   string // Pending, InProgress, Completed, Failed
	Result   string
}

func (t *Task) Process() error {
	rand := rand.Intn(2)
	Logger.Info().Msgf("Processing Task ID: %s, Name: %s", t.ID, t.Name)
	t.Status = "InProgress"
	time.Sleep(3 * time.Second)
	if rand == 1 {
		t.Status = "Completed"
		Logger.Info().Msgf("Task ID: %s completed successfully", t.ID)
		return nil
	} else {
		t.Status = "Failed"
		Logger.Error().Msgf("Task ID: %s failed during processing", t.ID)
		return errors.New("Task Failed")
	}
}
