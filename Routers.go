package main

import "github.com/gin-gonic/gin"

func SetupRoutes() {
	Router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to the Task Processor API")
	})
	Router.POST("/tasks", handleAddTask)
	Router.GET("/tasks", handleGetAllTasks)
	Router.GET("/task/:id", handleGetSpecificTask)
}
