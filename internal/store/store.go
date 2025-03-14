package store

import (
	"github.com/google/uuid"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
)

var TasksDB = make(map[uuid.UUID]entity.Task)
