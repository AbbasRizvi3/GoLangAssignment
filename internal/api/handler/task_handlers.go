package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/core/app"
	logger "github.com/AbbasRizvi3/GoLangAssignment.git/internal/logging"
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
	"github.com/gin-gonic/gin"
)

const (
	randomNumber = 1000000
)

func HandleAddTask(ctx *gin.Context) {

	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	temp := &tasks.Task{
		ID:       fmt.Sprintf("%d", rand.Intn(randomNumber)),
		Name:     req.Name,
		Priority: req.Priority,
		Status:   "Pending",
	}
	app.Tasks.AddTask(temp)
	temp.Mutex.RLock()
	defer temp.Mutex.RUnlock()
	taskCopy := copyTask(temp)
	if _, err := json.Marshal(gin.H{"message": "Task added successfully", "task": taskCopy}); err != nil {
		fmt.Printf("Failed to encode response JSON: %v\n", err)
		return
	}
	fmt.Printf("Added Task ID: %s, Name: %s\n", temp.ID, temp.Name)
	ctx.JSON(200, gin.H{"message": "Task added successfully", "task": taskCopy})
	app.TaskChannel <- struct{}{}
}

func copyTask(t *tasks.Task) tasks.TaskData {
	t.Mutex.RLock()
	defer t.Mutex.RUnlock()
	return tasks.TaskData{
		ID:       t.ID,
		Name:     t.Name,
		Priority: t.Priority,
		Status:   t.Status,
		Result:   t.Result,
	}
}

func HandleGetAllTasks(ctx *gin.Context) {
	fmt.Println("Fetching all tasks")
	app.Tasks.Mutex.RLock()
	defer app.Tasks.Mutex.RUnlock()

	tasksCopy := make([]tasks.Task, 0, len(app.Tasks.Tasks)+len(app.ResultSlice))
	for _, t := range app.Tasks.Tasks {
		if t != nil {
			taskCopy := copyTask(t)
			tasksCopy = append(tasksCopy, tasks.Task{
				ID:       taskCopy.ID,
				Name:     taskCopy.Name,
				Priority: taskCopy.Priority,
				Status:   taskCopy.Status,
				Result:   taskCopy.Result,
			})
		}
	}

	for i := range app.ResultSlice {
		t := &app.ResultSlice[i]
		tasksCopy = append(tasksCopy, tasks.Task{
			ID:       t.ID,
			Name:     t.Name,
			Priority: t.Priority,
			Status:   t.Status,
			Result:   t.Result,
		})
	}

	fmt.Printf("Total tasks fetched: %d\n", len(tasksCopy))
	if _, err := json.Marshal(tasksCopy); err != nil {
		logger.Logger.Error().Msgf("Failed to encode tasks: %v", err)
		return
	}
	ctx.JSON(200, gin.H{"tasks": tasksCopy})
}

func HandleGetSpecificTask(ctx *gin.Context) {

	app.Tasks.Mutex.RLock()
	defer app.Tasks.Mutex.RUnlock()
	id := ctx.Param("id")
	fmt.Printf("Fetching task with ID: %s\n", id)
	temp := make([]tasks.Task, 0)

	for _, t := range app.Tasks.Tasks {
		if t.ID == id {
			taskCopy := copyTask(t)
			temp = append(temp, tasks.Task{
				ID:       taskCopy.ID,
				Name:     taskCopy.Name,
				Priority: taskCopy.Priority,
				Status:   taskCopy.Status,
				Result:   taskCopy.Result,
			})
		}
	}

	if len(temp) == 0 {
		fmt.Printf("Task with ID: %s not found\n", id)
		if _, err := json.Marshal(gin.H{"error": "Task not found"}); err != nil {
			fmt.Printf("Failed to encode 404 JSON: %v\n", err)
			return
		}
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	fmt.Printf("Task with ID: %s found\n", id)
	if _, err := json.Marshal(gin.H{"tasks": temp}); err != nil {
		fmt.Printf("Failed to encode tasks JSON: %v\n", err)
	}
	ctx.JSON(200, gin.H{"tasks": temp})
}
