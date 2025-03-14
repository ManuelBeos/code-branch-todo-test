package mappers

import (
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
	"github.com/manuelbeos/code-branch-todo-test/internal/handlers/dtos"
)

func MapperUpdateTaskRequestToTaskEntity(updateReq dtos.UpdateTaskRequestDto) entity.Task {
	return entity.Task{
		Id:          updateReq.Id,
		Title:       updateReq.Title,
		Description: updateReq.Description,
		IsCompleted: updateReq.IsCompleted,
	}
}
