package scheduler

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func Gocron_scheduler() {
	// Do jobs without params
	gocron.Every(2).Seconds().Do(task)

	// Do jobs with params
	gocron.Every(1).Minute().Do(taskWithParams, 1, "hello")

	<-gocron.Start()

	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	//gocron.Every(1).Monday().At("18:30").Do(task)

	// Remove a specific job
	//gocron.Remove(task)

	// Clear all scheduled jobs
	//gocron.Clear()
}
