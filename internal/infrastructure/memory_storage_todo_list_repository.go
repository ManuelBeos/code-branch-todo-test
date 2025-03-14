package infrastructure

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
	domain "github.com/manuelbeos/code-branch-todo-test/internal/domain/errors"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/repository"
	"github.com/manuelbeos/code-branch-todo-test/internal/utils"
)

type MemoryStorageTodoListRepository struct {
	memoryTasks map[uuid.UUID]entity.Task
}

func NewMemoryStorageTodoListRepository(tasks map[uuid.UUID]entity.Task) repository.TodoListRepository {
	return &MemoryStorageTodoListRepository{memoryTasks: tasks}
}

func (mr *MemoryStorageTodoListRepository) CreateTask(ctx context.Context, newTask entity.Task) (*entity.Task, error) {
	chanResponse := make(chan entity.Task)

	go func() {

		sleepTime := time.Duration(utils.RandomNumber(500, 2000)) * time.Millisecond
		time.Sleep(sleepTime)

		chanResponse <- newTask
		close(chanResponse)
	}()

	taskCreated := <-chanResponse

	mr.memoryTasks[newTask.Id] = taskCreated

	return &newTask, nil
}

func (mr *MemoryStorageTodoListRepository) GetAllTasks(ctx context.Context) ([]*entity.Task, error) {
	chanResponse := make(chan []*entity.Task)

	go func() {
		tasks := make([]*entity.Task, 0, len(mr.memoryTasks))
		for _, task := range mr.memoryTasks {
			tasks = append(tasks, &task)
		}
		sleepTime := time.Duration(utils.RandomNumber(500, 2000)) * time.Millisecond
		time.Sleep(sleepTime)

		chanResponse <- tasks
		close(chanResponse)
	}()

	tasksResponse := <-chanResponse

	if len(tasksResponse) == 0 {
		return nil, domain.ErrThereAreNoTasks
	}

	return tasksResponse, nil
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
