package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
	domain "github.com/manuelbeos/code-branch-todo-test/internal/domain/errors"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/repository"
)

type MemoryStorageTodoListRepository struct {
	memoryTasks map[uuid.UUID]entity.Task
}

func NewMemoryStorageTodoListRepository(tasks map[uuid.UUID]entity.Task) repository.TodoListRepository {
	return &MemoryStorageTodoListRepository{memoryTasks: tasks}
}

func (mr *MemoryStorageTodoListRepository) CreateTask(ctx context.Context, newTask entity.Task) (*entity.Task, error) {
	mr.memoryTasks[newTask.Id] = newTask

	return &newTask, nil
}

func (mr *MemoryStorageTodoListRepository) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	tasks := make([]*entity.Task, 0, len(mr.memoryTasks))
	for _, task := range mr.memoryTasks {
		tasks = append(tasks, &task)
	}

	if len(tasks) == 0 {
		return nil, domain.ErrThereAreNoTasks
	}

	return tasks, nil
}

func (mr *MemoryStorageTodoListRepository) GetTaskByID(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	task, ok := mr.memoryTasks[id]
	if !ok {
		return nil, domain.ErrTaskNotFound
	}

	return &task, nil
}

func (mr *MemoryStorageTodoListRepository) UpdateTask(ctx context.Context, updatedTask *entity.Task) (*entity.Task, error) {
	mr.memoryTasks[updatedTask.Id] = *updatedTask
	return updatedTask, nil
}

func (mr *MemoryStorageTodoListRepository) DeleteTask(ctx context.Context, id uuid.UUID) error {
	delete(mr.memoryTasks, id)

	return nil
}
