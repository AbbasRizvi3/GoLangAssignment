package main

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

var Tasks []Task

func worker(jobs <-chan Task, results chan<- Task) {
	for j := range jobs {
		// fmt.Println("worker", id, "started  job", j)
		// time.Sleep(time.Second)
		// fmt.Println("worker", id, "finished job", j)
		// results <- j * 2
		j.Process() // deal with the error returned.
		results <- j
	}
}

func main() {

	const numJobs = 50
	taskChannel := make(chan Task, numJobs)
	resultChannel := make(chan Task, numJobs)

	r := gin.Default()

	for w := 0; w < 5; w++ {
		Logger.Info().Msgf("Starting Worker %d", w)
		go worker(taskChannel, resultChannel)
	}

	for j := 0; j < numJobs; j++ {
		taskChannel <- Task{
			ID:       fmt.Sprintf("%d", rand.Intn(1000)),
			Name:     "sample job",
			Priority: 1,
		}
	}
	close(taskChannel)

	for a := 0; a < numJobs; a++ {
		<-resultChannel
	}

	Logger.Info().Msg("Logger Exiting (Program execution suspended)")
}
