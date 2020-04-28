package main

import (
	"fmt"
	"time"
	"xcrap/api/queue"
)

// CreateJob  define your tasks here
func CreateJob(name string) queue.Job {
	// Create job
	j := queue.Job{}
	j.Name = name
	j.Task = queue.CreateTask(func() {
		time.Sleep(5000 * time.Millisecond)
		fmt.Print("Task ...1")
	})
	return j
}
