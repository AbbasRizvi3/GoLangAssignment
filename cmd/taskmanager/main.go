package main

import (
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/core/app"
	logger "github.com/AbbasRizvi3/GoLangAssignment.git/internal/logging"
	routers "github.com/AbbasRizvi3/GoLangAssignment.git/internal/router"
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
)

const (
	workerCount = 5
	port        = ":8000"
)

func main() {

	routers.SetupRoutes()

	for w := 0; w < workerCount; w++ {
		logger.Logger.Info().Msgf("Starting Worker %d", w)
		go tasks.Worker(&app.Tasks, app.TaskChannel, app.ResultChannel)
	}

	go func() {
		for res := range app.ResultChannel {
			logger.Logger.Info().Msgf("Result received for Task ID: %s, Status: %s", res.ID, res.Status)
		}
	}()
	err := app.Router.Run(port)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Failed to start server")
	}
	logger.Logger.Info().Msg("Logger Exiting (Program execution suspended)")

}
