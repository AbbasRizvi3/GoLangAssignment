package main

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func handleAddTask(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	temp := &Task{
		ID:       fmt.Sprintf("%d", rand.Intn(1000000)),
		Name:     req.Name,
		Priority: req.Priority,
		Status:   "Pending",
	}
	Tasks.AddTask(temp)
	Logger.Info().Msgf("Added Task ID: %s, Name: %s", temp.ID, temp.Name)
	ctx.JSON(200, gin.H{"message": "Task added successfully", "task": temp})
	select {
	case taskChannel <- struct{}{}:
	default:
	}
}

func handleGetAllTasks(ctx *gin.Context) {
	Logger.Info().Msg("Fetching all tasks")
	Tasks.mutex.RLock()
	defer Tasks.mutex.RUnlock()
	ctx.JSON(200, gin.H{"tasks": Tasks.Tasks})
}

func handleGetSpecificTask(ctx *gin.Context) {

	Tasks.mutex.RLock()
	defer Tasks.mutex.RUnlock()
	id := ctx.Param("id")
	Logger.Info().Msgf("Fetching task with ID: %s", id)
	temp := make([]Task, 0)
	for _, item := range Tasks.Tasks {
		if item.ID == id {
			temp = append(temp, *item)
		}
	}

	ctx.JSON(200, gin.H{"tasks": temp})
}
