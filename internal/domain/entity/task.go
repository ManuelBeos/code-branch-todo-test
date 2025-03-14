package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(title string, description string) Task {
	return Task{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) Update(title string, description string, isCompleted bool) {
	t.Title = title
	t.Description = description
	t.IsCompleted = isCompleted
	t.UpdatedAt = time.Now()
}
