package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
)

type TodoListRepository interface {
	CreateTask(context.Context, entity.Task) (*entity.Task, error)
	GetAllTasks(context.Context) ([]*entity.Task, error)
	GetTaskByID(context.Context, uuid.UUID) (*entity.Task, error)
	UpdateTask(context.Context, *entity.Task) (*entity.Task, error)
	DeleteTask(context.Context, uuid.UUID) error
}
