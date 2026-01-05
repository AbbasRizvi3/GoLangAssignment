package routers

import (
	"github.com/AbbasRizvi3/GoLangAssignment.git/api/handler"
	"github.com/AbbasRizvi3/GoLangAssignment.git/core/app"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	app.Router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to the Task Processor API")
	})
	app.Router.POST("/tasks", handler.HandleAddTask)
	app.Router.GET("/tasks", handler.HandleGetAllTasks)
	app.Router.GET("/task/:id", handler.HandleGetSpecificTask)
}
