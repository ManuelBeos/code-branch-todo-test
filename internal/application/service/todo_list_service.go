package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/repository"
)

type TodoListService struct {
	repository repository.TodoListRepository
}

func NewTodoListService(repository repository.TodoListRepository) *TodoListService {
	return &TodoListService{repository: repository}
}

func (tls *TodoListService) CreateTask(ctx context.Context, title string, description string) (*entity.Task, error) {
	task := entity.NewTask(title, description)

	return tls.repository.CreateTask(ctx, task)
}

func (tls *TodoListService) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	return tls.repository.GetAllTasks(ctx)
}

func (tls *TodoListService) GetTaskByID(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	return tls.repository.GetTaskByID(ctx, id)
}

func (tls *TodoListService) UpdateTask(ctx context.Context, taskToUpdate entity.Task) (*entity.Task, error) {
	task, err := tls.repository.GetTaskByID(ctx, taskToUpdate.Id)

	if err != nil {
		return nil, err
	}

	task.Update(taskToUpdate.Title, taskToUpdate.Description, taskToUpdate.IsCompleted)

	return tls.repository.UpdateTask(ctx, task)
}

func (tls *TodoListService) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, err := tls.repository.GetTaskByID(ctx, id)
	if err != nil {
		return err
	}

	return tls.repository.DeleteTask(ctx, id)
}
