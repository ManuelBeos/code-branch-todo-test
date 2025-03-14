package error_response

import (
	"net/http"

	"github.com/manuelbeos/code-branch-todo-test/internal/handlers/dtos"
)

var (
	ErrReadingRequestBody = dtos.NewErrorResponse("Error reading request body", http.StatusBadRequest)
	ErrParsingRequestBody = dtos.NewErrorResponse("Error parsing request body", http.StatusBadRequest)
	ErrCreatingTask       = dtos.NewErrorResponse("Error creating task", http.StatusInternalServerError)
	ErrGettingTasks       = dtos.NewErrorResponse("Error getting all tasks", http.StatusInternalServerError)
	ErrParsingTaskID      = dtos.NewErrorResponse("Error parsing task id is not a valid uuid", http.StatusBadRequest)
	ErrGettingTaskByID    = dtos.NewErrorResponse("Error getting task by id", http.StatusInternalServerError)
	ErrUpdatingTask       = dtos.NewErrorResponse("Error updating task", http.StatusInternalServerError)
	ErrDeletingTask       = dtos.NewErrorResponse("Error deleting task", http.StatusInternalServerError)
	ErrThereAreNoTasks    = dtos.NewErrorResponse("There are no tasks", http.StatusNotFound)
	ErrTaskNotFound       = dtos.NewErrorResponse("Task not found", http.StatusNotFound)
)

//params

var (
	ErrTitleFieldIsRequired = dtos.NewErrorResponse("Title field is required", http.StatusBadRequest)
)
