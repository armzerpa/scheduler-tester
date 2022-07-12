package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/procyon-projects/chrono"
)

func Chrono_scheduler() {
	taskScheduler := chrono.NewDefaultTaskScheduler()

	task, err := taskScheduler.ScheduleAtFixedRate(func(ctx context.Context) {
		log.Print("Fixed Rate of 5 seconds")
	}, 5*time.Second)

	if err == nil {
		log.Print("Task has been scheduled successfully.")
	}

	count := 0
	for !task.IsCancelled() {
		time.Sleep(time.Second)
		count++
		if count == 100 {
			task.Cancel()
		}
	}
}

func Chrono_cron() {
	taskScheduler := chrono.NewDefaultTaskScheduler()

	task, err := taskScheduler.ScheduleWithCron(func(ctx context.Context) {
		log.Print("Scheduled Task With Cron")
	}, "0 45 18 10 * *")

	if err == nil {
		log.Print("Task has been scheduled")
	}

	count := 0
	for !task.IsCancelled() {
		time.Sleep(time.Second)
		count++
		if count == 100 {
			task.Cancel()
		}
	}
}
