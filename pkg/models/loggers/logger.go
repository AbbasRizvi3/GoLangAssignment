package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func init() {
	file, err := os.OpenFile("Tasks.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	Logger = zerolog.New(file).With().Timestamp().Logger()
	Logger.Info().Msg("Logger Started")
}
