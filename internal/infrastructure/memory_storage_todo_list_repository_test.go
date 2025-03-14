package infrastructure

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
	domain "github.com/manuelbeos/code-branch-todo-test/internal/domain/errors"
	"github.com/stretchr/testify/assert"
)

func getTestMemory() map[uuid.UUID]entity.Task {
	id := uuid.New()
	return map[uuid.UUID]entity.Task{
		id: {
			Id:          id,
			Title:       "Test",
			Description: "Test",
		},
	}
}

func TestNewMemoryStorageTodoListRepository_Success(t *testing.T) {
	memory := getTestMemory()
	asserts := assert.New(t)
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)
}

func TestMemoryStorageTodoListRepository_CreateTask_Success(t *testing.T) {
	asserts := assert.New(t)
	memory := getTestMemory()
	ctx := context.Background()
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)

	task := entity.Task{
		Id:          uuid.New(),
		Title:       "Test",
		Description: "Test",
	}

	taskCreated, err := memoryRepo.CreateTask(ctx, task)
	asserts.Nil(err)
	asserts.NotNil(taskCreated)
	asserts.Equal(task.Id, taskCreated.Id)
	asserts.Equal(task.Title, taskCreated.Title)
	asserts.Equal(task.Description, taskCreated.Description)
}

func TestMemoryStorageTodoListRepository_GetAllTasks_Success(t *testing.T) {
	asserts := assert.New(t)
	memory := getTestMemory()
	ctx := context.Background()
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)

	tasks, err := memoryRepo.GetAllTasks(ctx)
	asserts.Nil(err)
	asserts.NotNil(tasks)
	asserts.Equal(1, len(tasks))
}

func TestMemoryStorageTodoListRepository_GetAllTasks_Empty(t *testing.T) {
	asserts := assert.New(t)
	memory := make(map[uuid.UUID]entity.Task)
	ctx := context.Background()
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)

	tasks, err := memoryRepo.GetAllTasks(ctx)
	asserts.NotNil(err)
	asserts.Nil(tasks)
	asserts.ErrorIs(err, domain.ErrThereAreNoTasks)
}

func TestMemoryStorageTodoListRepository_GetTaskByID_Success(t *testing.T) {
	asserts := assert.New(t)
	memory := getTestMemory()
	ctx := context.Background()
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)

	task := entity.Task{
		Id:          uuid.New(),
		Title:       "Test",
		Description: "Test",
	}

	taskCreated, err := memoryRepo.CreateTask(ctx, task)
	asserts.Nil(err)
	asserts.NotNil(taskCreated)

	taskByID, err := memoryRepo.GetTaskByID(ctx, task.Id)
	asserts.Nil(err)
	asserts.NotNil(taskByID)
	asserts.Equal(task.Id, taskByID.Id)
	asserts.Equal(task.Title, taskByID.Title)
	asserts.Equal(task.Description, taskByID.Description)
}

func TestMemoryStorageTodoListRepository_GetTaskByID_NotFound(t *testing.T) {
	asserts := assert.New(t)
	memory := getTestMemory()
	ctx := context.Background()
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)

	taskByID, err := memoryRepo.GetTaskByID(ctx, uuid.New())
	asserts.NotNil(err)
	asserts.Nil(taskByID)
	asserts.ErrorIs(err, domain.ErrTaskNotFound)
}

func TestMemoryStorageTodoListRepository_DeleteTask_Success(t *testing.T) {
	asserts := assert.New(t)
	memory := getTestMemory()
	ctx := context.Background()
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)

	task := entity.Task{
		Id:          uuid.New(),
		Title:       "Test",
		Description: "Test",
	}

	taskCreated, err := memoryRepo.CreateTask(ctx, task)
	asserts.Nil(err)
	asserts.NotNil(taskCreated)
	asserts.Equal(2, len(memory))

	err = memoryRepo.DeleteTask(ctx, task.Id)
	asserts.Nil(err)

	taskByID, err := memoryRepo.GetTaskByID(ctx, task.Id)
	asserts.NotNil(err)
	asserts.Nil(taskByID)
	asserts.ErrorIs(err, domain.ErrTaskNotFound)
}

func TestMemoryStorageTodoListRepository_UpdateTask_Success(t *testing.T) {
	asserts := assert.New(t)
	memory := getTestMemory()
	ctx := context.Background()
	memoryRepo := NewMemoryStorageTodoListRepository(memory)
	asserts.NotNil(memoryRepo)

	task := entity.Task{
		Id:          uuid.New(),
		Title:       "Test",
		Description: "Test",
	}

	taskCreated, err := memoryRepo.CreateTask(ctx, task)
	asserts.Nil(err)
	asserts.NotNil(taskCreated)

	taskUpdated := &entity.Task{
		Id:          task.Id,
		Title:       "Test Updated",
		Description: "Test Updated",
	}

	taskUpdated, err = memoryRepo.UpdateTask(ctx, taskUpdated)
	asserts.Nil(err)
	asserts.NotNil(taskUpdated)

	taskByID, err := memoryRepo.GetTaskByID(ctx, task.Id)
	asserts.Nil(err)
	asserts.NotNil(taskByID)
	asserts.Equal(taskUpdated.Id, taskByID.Id)
	asserts.Equal(taskUpdated.Title, taskByID.Title)
	asserts.Equal(taskUpdated.Description, taskByID.Description)
}
