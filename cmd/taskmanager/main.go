package main

import (
	"fmt"
	"sync"

	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/core/app"
	routers "github.com/AbbasRizvi3/GoLangAssignment.git/internal/router"
	"github.com/AbbasRizvi3/GoLangAssignment.git/internal/tasks"
)

const (
	workerCount = 5
	port        = ":8000"
)

var mutex sync.Mutex
var ActiveWorkers = 0

func incrementActiveWorkers() {
	mutex.Lock()
	defer mutex.Unlock()
	ActiveWorkers++
}

func main() {

	routers.SetupRoutes()

	go func() {
		for range app.TaskChannel {
			fmt.Println("Task received in TaskChannel")
			if ActiveWorkers < workerCount {
				incrementActiveWorkers()
				fmt.Printf("Active Workers: %d\n", ActiveWorkers)
				go tasks.Worker(&app.Tasks, &app.ResultSlice, &ActiveWorkers, &mutex)
			}
		}
	}()

	err := app.Router.Run(port)
	if err != nil {
		fmt.Println("Failed to start server")
	}
	fmt.Println("Logger Exiting (Program execution suspended")

}
