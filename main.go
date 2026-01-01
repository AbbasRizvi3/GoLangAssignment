package main

import (
	"github.com/gin-gonic/gin"
)

var Tasks []*Task

var taskChannel = make(chan *Task, 100)
var resultChannel = make(chan *Task, 100)
var Router = gin.Default()

// var numJobs = 50

func main() {

	SetupRoutes()

	for w := 0; w < 5; w++ {
		Logger.Info().Msgf("Starting Worker %d", w)
		go worker(taskChannel, resultChannel)
	}

	go func() {
		for res := range resultChannel {
			Logger.Info().Msgf("Result received for Task ID: %s, Status: %s", res.ID, res.Status)
		}
	}()
	Router.Run(":8000")
	Logger.Info().Msg("Logger Exiting (Program execution suspended)")

}
