package types

import (
	"time"
)

// Picon a Pipeline In Containers
type Picon struct {
	Name    string
	Workers []Worker
	Stages  []Stage
}

// Worker a worker to be used in the pipeline via tasks
type Worker struct {
	ID    string
	Image string
}

// Stage to define big part of the pipeline
type Stage struct {
	// Name represents the human-readable name of the stage
	Name string

	Tasks []Task
}

// Task to be executed by a worker
type Task struct {
	Name     string
	WorkerID string
	Params   string
	Type     string
	Timeout  time.Duration
}
