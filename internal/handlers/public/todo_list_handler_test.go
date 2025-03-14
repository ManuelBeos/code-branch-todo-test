package public

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manuelbeos/code-branch-todo-test/internal/application/service"
	"github.com/manuelbeos/code-branch-todo-test/internal/domain/entity"
	domain "github.com/manuelbeos/code-branch-todo-test/internal/domain/errors"
	"github.com/manuelbeos/code-branch-todo-test/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTodoListHandler_Success(t *testing.T) {
	asserts := assert.New(t)
	mockRepo := mocks.NewTodoListRepository(t)
	service := service.NewTodoListService(mockRepo)
	handler := NewTodoListHandler(service)
	asserts.NotNil(handler)
}

func TestTodoListHandler_CreateNewTask(t *testing.T) {
	asserts := assert.New(t)
	mockError := errors.New(" mockerror")

	task := entity.NewTask("title", "description")

	tests := []struct {
		name                    string
		body                    string
		expectedStatusCode      int
		repoCreateTaskError     error
		setCustomReturnMockRepo bool
		expectedResponse        string
		validateBodyResponse    bool
		repoTask                *entity.Task
	}{
		{
			name:                    "CreateNewTask - Success",
			body:                    `{"title": "title", "description": "description"}`,
			expectedStatusCode:      http.StatusCreated,
			repoCreateTaskError:     nil,
			setCustomReturnMockRepo: true,
			validateBodyResponse:    false,
			repoTask:                &task,
		},
		{
			name:                    "CreateNewTask - Error Creating task",
			body:                    `{"title": "title", "description": "description"}`,
			expectedStatusCode:      http.StatusInternalServerError,
			repoCreateTaskError:     mockError,
			setCustomReturnMockRepo: true,
			expectedResponse:        `{"message":"Error creating task","code":500}`,
			validateBodyResponse:    true,
			repoTask:                nil,
		},
		{
			name:                    "CreateNewTask - Error title field empty",
			body:                    `{"title": "", "description": "description"}`,
			expectedStatusCode:      http.StatusBadRequest,
			repoCreateTaskError:     nil,
			setCustomReturnMockRepo: false,
			expectedResponse:        `{"message":"Title field is required","code":400}`,
			validateBodyResponse:    true,
		},
		{
			name:                    "CreateNewTask - Error unmarshal body",
			body:                    `"title": "", "description": "description"}`,
			expectedStatusCode:      http.StatusBadRequest,
			repoCreateTaskError:     nil,
			setCustomReturnMockRepo: false,
			expectedResponse:        `{"message":"Error parsing request body","code":400}`,
			validateBodyResponse:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewTodoListRepository(t)

			var bodyRequest io.ReadCloser

			buffer := new(bytes.Buffer)
			buffer.Write([]byte(tt.body))
			bodyRequest = io.NopCloser(buffer)

			ctx := context.Background()

			if tt.setCustomReturnMockRepo {
				mockRepo.On("CreateTask", ctx, mock.Anything).Return(tt.repoTask, tt.repoCreateTaskError)
			}

			service := service.NewTodoListService(mockRepo)
			muxRouter := mux.NewRouter()
			handler := NewTodoListHandler(service)
			handler.RegisterEndpoints(muxRouter)

			req := httptest.NewRequest(http.MethodPost, "/tasks", bodyRequest)
			w := httptest.NewRecorder()

			handler.CreateNewTask(w, req.WithContext(ctx))

			asserts.Equal(tt.expectedStatusCode, w.Code)

			if tt.validateBodyResponse {
				asserts.Equal(tt.expectedResponse, w.Body.String())
			}
		})
	}

}

func TestTodoListHandler_UpdateTask(t *testing.T) {
	asserts := assert.New(t)
	mockError := errors.New("error mockerror")
	task := entity.NewTask("title", "description")

	tests := []struct {
		name                    string
		body                    string
		expectedStatusCode      int
		repoUpdateTaskError     error
		repoGetTaskError        error
		setCustomReturnMockRepo bool
		expectedResponse        string
		validateBodyResponse    bool
		repoTask                *entity.Task
		taskIdUrl               string
	}{
		{
			name:                    "UpdateTask - Success",
			body:                    `{"title": "title", "description": "description", "is_completed": true}`,
			expectedStatusCode:      http.StatusOK,
			repoUpdateTaskError:     nil,
			repoGetTaskError:        nil,
			setCustomReturnMockRepo: true,
			validateBodyResponse:    false,
			repoTask:                &task,
			taskIdUrl:               task.Id.String(),
		},
		{
			name:                    "UpdateTask - Error Updating task",
			body:                    `{"title": "title", "description": "description", "is_completed": true}`,
			expectedStatusCode:      http.StatusInternalServerError,
			repoUpdateTaskError:     mockError,
			repoGetTaskError:        nil,
			setCustomReturnMockRepo: true,
			validateBodyResponse:    true,
			repoTask:                &task,
			taskIdUrl:               task.Id.String(),
			expectedResponse:        `{"message":"Error updating task","code":500}`,
		},
		{
			name:                    "UpdateTask - Error Updating task not found",
			body:                    `{"title": "title", "description": "description", "is_completed": true}`,
			expectedStatusCode:      http.StatusNotFound,
			repoUpdateTaskError:     domain.ErrTaskNotFound,
			repoGetTaskError:        nil,
			setCustomReturnMockRepo: true,
			validateBodyResponse:    true,
			repoTask:                &task,
			taskIdUrl:               task.Id.String(),
			expectedResponse:        `{"message":"Task not found","code":404}`,
		},
		{
			name:                    "UpdateTask - Error validating fields title empty",
			body:                    `{ "description": "description", "is_completed": true}`,
			expectedStatusCode:      http.StatusBadRequest,
			repoUpdateTaskError:     nil,
			repoGetTaskError:        nil,
			setCustomReturnMockRepo: false,
			validateBodyResponse:    true,
			repoTask:                &task,
			taskIdUrl:               task.Id.String(),
			expectedResponse:        `{"message":"Title field is required","code":400}`,
		},
		{
			name:                    "UpdateTask - Error unmarshal body",
			body:                    `bad`,
			expectedStatusCode:      http.StatusBadRequest,
			repoUpdateTaskError:     nil,
			repoGetTaskError:        nil,
			setCustomReturnMockRepo: false,
			validateBodyResponse:    true,
			repoTask:                &task,
			taskIdUrl:               task.Id.String(),
			expectedResponse:        `{"message":"Error parsing request body","code":400}`,
		},
		{
			name:                    "UpdateTask - Error parsing id from url",
			body:                    `{"title": "title", "description": "description", "is_completed": true}`,
			expectedStatusCode:      http.StatusBadRequest,
			repoUpdateTaskError:     nil,
			repoGetTaskError:        nil,
			setCustomReturnMockRepo: false,
			validateBodyResponse:    true,
			repoTask:                &task,
			taskIdUrl:               "bad-id",
			expectedResponse:        `{"message":"Error parsing task id is not a valid uuid","code":400}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewTodoListRepository(t)

			var bodyRequest io.ReadCloser

			buffer := new(bytes.Buffer)
			buffer.Write([]byte(tt.body))
			bodyRequest = io.NopCloser(buffer)

			ctx := context.Background()

			service := service.NewTodoListService(mockRepo)
			muxRouter := mux.NewRouter()
			handler := NewTodoListHandler(service)
			handler.RegisterEndpoints(muxRouter)

			req := httptest.NewRequest(http.MethodPut, "/tasks/{id}", bodyRequest)
			req = req.WithContext(ctx)
			req = mux.SetURLVars(req, map[string]string{"id": tt.taskIdUrl})

			if tt.setCustomReturnMockRepo {
				mockRepo.On("GetTaskByID", req.Context(), mock.Anything).Return(tt.repoTask, tt.repoGetTaskError)
				mockRepo.On("UpdateTask", req.Context(), mock.Anything).Return(tt.repoTask, tt.repoUpdateTaskError)
			}

			w := httptest.NewRecorder()

			handler.UpdateTask(w, req)

			asserts.Equal(tt.expectedStatusCode, w.Code)

			if tt.validateBodyResponse {
				asserts.Equal(tt.expectedResponse, w.Body.String())
			}
		})
	}

}
