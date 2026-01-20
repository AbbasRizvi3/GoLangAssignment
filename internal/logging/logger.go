package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

const (
	logFileName    = "tasks.log"
	fileAccessMode = os.O_APPEND | os.O_CREATE | os.O_WRONLY
)

func init() {
	file, err := os.OpenFile(logFileName, fileAccessMode, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	Logger = zerolog.New(file).With().Timestamp().Logger()
	fmt.Println("Logger Started")
}
