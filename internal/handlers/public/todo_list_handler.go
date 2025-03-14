package public

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/manuelbeos/code-branch-todo-test/internal/application/service"
	domain "github.com/manuelbeos/code-branch-todo-test/internal/domain/errors"
	"github.com/manuelbeos/code-branch-todo-test/internal/handlers/dtos"
	error_response "github.com/manuelbeos/code-branch-todo-test/internal/handlers/errors"
	"github.com/manuelbeos/code-branch-todo-test/internal/handlers/mappers"
	handler_utils "github.com/manuelbeos/code-branch-todo-test/internal/handlers/utils"
)

type TodoListHandler struct {
	service *service.TodoListService
}

func NewTodoListHandler(service *service.TodoListService) *TodoListHandler {
	return &TodoListHandler{service: service}
}

func (tlh *TodoListHandler) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrReadingRequestBody)
		return
	}

	createNewTaskReq := &dtos.CreateTaskRequestDto{}
	err = json.Unmarshal(body, createNewTaskReq)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrParsingRequestBody)
		return
	}

	ok := createNewTaskReq.ValidTitleField()

	if !ok {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrTitleFieldIsRequired)
		return
	}

	task, err := tlh.service.CreateTask(ctx, createNewTaskReq.Title, createNewTaskReq.Description)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusInternalServerError, error_response.ErrCreatingTask)
		return
	}

	handler_utils.HandlerSuccessResponse(w, http.StatusCreated, task)
}

func (tlh *TodoListHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tasks, err := tlh.service.GetAllTasks(ctx)
	if err != nil {

		if errors.Is(err, domain.ErrThereAreNoTasks) {
			handler_utils.HandlerErrorResponse(w, http.StatusNotFound, error_response.ErrThereAreNoTasks)
			return
		}

		handler_utils.HandlerErrorResponse(w, http.StatusInternalServerError, error_response.ErrGettingTasks)
		return
	}

	handler_utils.HandlerSuccessResponse(w, http.StatusOK, tasks)
}

// GetTaskByID retrieves a task by its ID.
// @Summary Get a task by ID
// @Description Get a task by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} entity.Task
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /tasks/{id} [get]
func (tlh *TodoListHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID := mux.Vars(r)["id"]
	taskIdAsUUID, err := uuid.Parse(taskID)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrParsingTaskID)
		return
	}

	task, err := tlh.service.GetTaskByID(ctx, taskIdAsUUID)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			handler_utils.HandlerErrorResponse(w, http.StatusNotFound, error_response.ErrTaskNotFound)
			return
		}

		handler_utils.HandlerErrorResponse(w, http.StatusInternalServerError, error_response.ErrGettingTaskByID)
		return
	}

	handler_utils.HandlerSuccessResponse(w, http.StatusOK, task)
}

func (tlh *TodoListHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID := mux.Vars(r)["id"]
	taskIdAsUUID, err := uuid.Parse(taskID)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrParsingTaskID)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrReadingRequestBody)
		return
	}

	updateTaskReq := &dtos.UpdateTaskRequestDto{}
	err = json.Unmarshal(body, updateTaskReq)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrParsingRequestBody)
		return
	}

	ok := updateTaskReq.ValidTitleField()

	if !ok {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrTitleFieldIsRequired)
		return
	}

	updateTaskReq.Id = taskIdAsUUID
	taskToUpdate := mappers.MapperUpdateTaskRequestToTaskEntity(*updateTaskReq)

	task, err := tlh.service.UpdateTask(ctx, taskToUpdate)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			handler_utils.HandlerErrorResponse(w, http.StatusNotFound, error_response.ErrTaskNotFound)
			return
		}

		handler_utils.HandlerErrorResponse(w, http.StatusInternalServerError, error_response.ErrUpdatingTask)
		return
	}

	handler_utils.HandlerSuccessResponse(w, http.StatusOK, task)

}

func (tlh *TodoListHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID := mux.Vars(r)["id"]
	taskIdAsUUID, err := uuid.Parse(taskID)
	if err != nil {
		handler_utils.HandlerErrorResponse(w, http.StatusBadRequest, error_response.ErrParsingTaskID)
		return
	}

	err = tlh.service.DeleteTask(ctx, taskIdAsUUID)
	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
			handler_utils.HandlerErrorResponse(w, http.StatusNotFound, error_response.ErrTaskNotFound)
			return
		}

		handler_utils.HandlerErrorResponse(w, http.StatusInternalServerError, error_response.ErrDeletingTask)
		return
	}

	handler_utils.HandlerSuccessResponse(w, http.StatusNoContent, nil)
}

func (tlh *TodoListHandler) RegisterEndpoints(r *mux.Router) {
	r.HandleFunc("/tasks", tlh.CreateNewTask).Methods(http.MethodPost)
	r.HandleFunc("/tasks", tlh.GetAllTasks).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", tlh.GetTaskByID).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", tlh.UpdateTask).Methods(http.MethodPut)
	r.HandleFunc("/tasks/{id}", tlh.DeleteTask).Methods(http.MethodDelete)
}
