package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/manuelbeos/code-branch-todo-test/docs" // docs is generated by Swaggo
	"github.com/manuelbeos/code-branch-todo-test/internal/application/service"
	"github.com/manuelbeos/code-branch-todo-test/internal/handlers/middlewares"
	"github.com/manuelbeos/code-branch-todo-test/internal/handlers/public"
	"github.com/manuelbeos/code-branch-todo-test/internal/infrastructure"
	"github.com/manuelbeos/code-branch-todo-test/internal/store"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer() *Server {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("¡Server up!"))
	})

	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &Server{httpServer: srv, router: r}
}

func (s *Server) Run() error {

	// dependency injection
	memoryStorageRepo := infrastructure.NewMemoryStorageTodoListRepository(store.TasksDB)
	todoListService := service.NewTodoListService(memoryStorageRepo)

	// handlers
	public.NewTodoListHandler(todoListService).RegisterEndpoints(s.router)

	//middlewares
	s.router.Use(middlewares.LoggingMiddleware)

	s.router.Use(mux.CORSMethodMiddleware(s.router))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Server initialized on http://localhost:8080")
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error trying to start the server: %v", err)
		}
	}()

	<-stop
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Error trying to shutdown the server: %v", err)
	}

	log.Println("Server stopped")

	return nil
}
