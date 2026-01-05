package handler

import (
	"fmt"
	"math/rand"

	"github.com/AbbasRizvi3/GoLangAssignment.git/core/app"
	"github.com/AbbasRizvi3/GoLangAssignment.git/core/workers"
	logger "github.com/AbbasRizvi3/GoLangAssignment.git/pkg/models/loggers"
	request "github.com/AbbasRizvi3/GoLangAssignment.git/pkg/models/req"
	"github.com/gin-gonic/gin"
)

func HandleAddTask(ctx *gin.Context) {
	var req request.Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	temp := &workers.Task{
		ID:       fmt.Sprintf("%d", rand.Intn(1000000)),
		Name:     req.Name,
		Priority: req.Priority,
		Status:   "Pending",
	}
	app.Tasks.AddTask(temp)
	logger.Logger.Info().Msgf("Added Task ID: %s, Name: %s", temp.ID, temp.Name)
	ctx.JSON(200, gin.H{"message": "Task added successfully", "task": temp})
	select {
	case app.TaskChannel <- struct{}{}:
	default:
	}
}

func HandleGetAllTasks(ctx *gin.Context) {
	logger.Logger.Info().Msg("Fetching all tasks")
	app.Tasks.Mutex.RLock()
	defer app.Tasks.Mutex.RUnlock()
	ctx.JSON(200, gin.H{"tasks": app.Tasks.Tasks})
}

func HandleGetSpecificTask(ctx *gin.Context) {

	app.Tasks.Mutex.RLock()
	defer app.Tasks.Mutex.RUnlock()
	id := ctx.Param("id")
	logger.Logger.Info().Msgf("Fetching task with ID: %s", id)
	temp := make([]workers.Task, 0)
	for _, item := range app.Tasks.Tasks {
		if item.ID == id {
			temp = append(temp, *item)
		}
	}

	ctx.JSON(200, gin.H{"tasks": temp})
}
