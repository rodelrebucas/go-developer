package queue

// Tasker interface
type Tasker interface {
	Execute()
}

// Task type
type Task struct {
	T func()
}

// Execute task
func (T *Task) Execute() {
	T.T()
}

// CreateTask task
func CreateTask(f func()) Tasker {
	return &Task{f}
}
