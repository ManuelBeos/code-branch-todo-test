package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
	"github.com/manuelbeos/code-branch-todo-test/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTodoListService_Success(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	service := NewTodoListService(mockRepository)

	asserts.NotNil(service)
}

func TestTodoListService_CreateTask_Success(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	ctx := context.Background()
	mockRepository.On("CreateTask", ctx, mock.Anything).Return(nil, nil)
	service := NewTodoListService(mockRepository)

	_, err := service.CreateTask(ctx, "title", "description")

	asserts.Nil(err)
}

func TestTodoListService_CreateTask_Error(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	mockError := errors.New("mock error")
	ctx := context.Background()
	mockRepository.On("CreateTask", ctx, mock.Anything).Return(nil, mockError)
	service := NewTodoListService(mockRepository)

	_, err := service.CreateTask(ctx, "title", "description")

	asserts.ErrorIs(mockError, err)
}

func TestTodoListService_GetAllTasks_Success(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	ctx := context.Background()
	mockRepository.On("GetAllTasks", ctx).Return(nil, nil)
	service := NewTodoListService(mockRepository)

	_, err := service.GetAllTasks(ctx)

	asserts.Nil(err)
}

func TestTodoListService_GetAllTasks_Error(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	mockError := errors.New("mock error")
	ctx := context.Background()
	mockRepository.On("GetAllTasks", ctx).Return(nil, mockError)
	service := NewTodoListService(mockRepository)

	_, err := service.GetAllTasks(ctx)

	asserts.ErrorIs(mockError, err)
}

func TestTodoListService_GetTaskByID_Success(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	ctx := context.Background()
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(nil, nil)
	service := NewTodoListService(mockRepository)

	_, err := service.GetTaskByID(ctx, uuid.New())

	asserts.Nil(err)
}

func TestTodoListService_GetTaskByID_Error(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	mockError := errors.New("mock error")
	ctx := context.Background()
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(nil, mockError)
	service := NewTodoListService(mockRepository)

	_, err := service.GetTaskByID(ctx, uuid.New())

	asserts.ErrorIs(mockError, err)
}

func TestTodoListService_UpdateTask_Success(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	ctx := context.Background()
	task := entity.Task{Id: uuid.New(), Title: "title", Description: "description", IsCompleted: false}
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(&task, nil)
	mockRepository.On("UpdateTask", ctx, mock.Anything).Return(nil, nil)
	service := NewTodoListService(mockRepository)

	_, err := service.UpdateTask(ctx, task)

	asserts.Nil(err)
}

func TestTodoListService_UpdateTask_Error_Not_Found(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	mockError := errors.New("mock error")
	ctx := context.Background()
	task := entity.Task{Id: uuid.New(), Title: "title", Description: "description", IsCompleted: false}
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(nil, mockError)
	service := NewTodoListService(mockRepository)

	_, err := service.UpdateTask(ctx, task)

	asserts.ErrorIs(mockError, err)
}

func TestTodoListService_UpdateTask_Error(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	mockError := errors.New("mock error")
	ctx := context.Background()
	task := entity.Task{Id: uuid.New(), Title: "title", Description: "description", IsCompleted: false}
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(&task, nil)
	mockRepository.On("UpdateTask", ctx, mock.Anything).Return(nil, mockError)
	service := NewTodoListService(mockRepository)

	_, err := service.UpdateTask(ctx, task)

	asserts.ErrorIs(mockError, err)
}

func TestTodoListService_DeleteTask_Success(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	ctx := context.Background()
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(nil, nil)
	mockRepository.On("DeleteTask", ctx, mock.Anything).Return(nil)
	service := NewTodoListService(mockRepository)

	err := service.DeleteTask(ctx, uuid.New())

	asserts.Nil(err)
}

func TestTodoListService_DeleteTask_Error_Not_Found(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	mockError := errors.New("mock error")
	ctx := context.Background()
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(nil, mockError)
	service := NewTodoListService(mockRepository)

	err := service.DeleteTask(ctx, uuid.New())

	asserts.ErrorIs(mockError, err)
}

func TestTodoListService_DeleteTask_Error(t *testing.T) {
	asserts := assert.New(t)
	mockRepository := mocks.NewTodoListRepository(t)
	mockError := errors.New("mock error")
	ctx := context.Background()
	mockRepository.On("GetTaskByID", ctx, mock.Anything).Return(nil, nil)
	mockRepository.On("DeleteTask", ctx, mock.Anything).Return(mockError)
	service := NewTodoListService(mockRepository)

	err := service.DeleteTask(ctx, uuid.New())

	asserts.ErrorIs(mockError, err)
}
