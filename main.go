package main

import (
	"fmt"

	"github.com/armzerpa/scheduler-tester/scheduler"
)

func main() {
	fmt.Println("scheduler test")
	//scheduler.Chrono_scheduler()
	scheduler.Gocron_scheduler()
	//scheduler.Chrono_cron()
}
