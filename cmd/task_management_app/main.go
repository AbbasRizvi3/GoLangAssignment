package main

import (
	"github.com/AbbasRizvi3/GoLangAssignment.git/core/app"
	"github.com/AbbasRizvi3/GoLangAssignment.git/core/workers"
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/routers"
	logger "github.com/AbbasRizvi3/GoLangAssignment.git/pkg/models/loggers"
)

func main() {

	routers.SetupRoutes()

	for w := 0; w < 5; w++ {
		logger.Logger.Info().Msgf("Starting Worker %d", w)
		go workers.Worker(&app.Tasks, app.TaskChannel, app.ResultChannel)
	}

	go func() {
		for res := range app.ResultChannel {
			logger.Logger.Info().Msgf("Result received for Task ID: %s, Status: %s", res.ID, res.Status)
		}
	}()
	app.Router.Run(":8000")
	logger.Logger.Info().Msg("Logger Exiting (Program execution suspended)")

}
