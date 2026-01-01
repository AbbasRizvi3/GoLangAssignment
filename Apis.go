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
	Tasks = append(Tasks, temp)
	Logger.Info().Msgf("Added Task ID: %s, Name: %s", temp.ID, temp.Name)
	ctx.JSON(200, gin.H{"message": "Task added successfully", "task": temp})
	taskChannel <- temp
}

func handleGetAllTasks(ctx *gin.Context) {
	Logger.Info().Msg("Fetching all tasks")
	ctx.JSON(200, gin.H{"tasks": Tasks})
}

func handleGetSpecificTask(ctx *gin.Context) {
	id := ctx.Param("id")
	temp := make([]Task, 0)
	for _, item := range Tasks {
		if item.ID == id {
			temp = append(temp, *item)
		}
	}
	Logger.Info().Msgf("Fetching task with ID: %s", id)
	ctx.JSON(200, gin.H{"tasks": temp})
}
