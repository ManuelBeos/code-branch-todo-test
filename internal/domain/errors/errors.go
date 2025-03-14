package domain

import "errors"

var (
	ErrTaskNotFound    = errors.New("task not found")
	ErrThereAreNoTasks = errors.New("there are no tasks created yet")
)
