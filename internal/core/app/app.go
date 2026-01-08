package app

import (
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
	"github.com/gin-gonic/gin"
)

const (
	taskChannelBufferSize   = 15
	resultChannelBufferSize = 100
)

var Tasks tasks.TaskQueue

var TaskChannel = make(chan struct{}, taskChannelBufferSize)
var ResultChannel = make(chan *tasks.Task, resultChannelBufferSize)
var Router = gin.Default()
