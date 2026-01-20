package app

import (
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
	"github.com/gin-gonic/gin"
)

const (
	taskChannelBufferSize = 30
)

var Tasks tasks.TaskQueue

var TaskChannel = make(chan struct{}, taskChannelBufferSize)
var ResultSlice []*tasks.Task
var Router = gin.Default()
