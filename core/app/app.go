package app

import (
	"github.com/AbbasRizvi3/GoLangAssignment.git/core/workers"
	"github.com/gin-gonic/gin"
)

var Tasks workers.TaskQueue

var TaskChannel = make(chan struct{}, 1)
var ResultChannel = make(chan *workers.Task, 100)
var Router = gin.Default()
