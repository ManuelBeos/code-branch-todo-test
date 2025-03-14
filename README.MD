# Todo List Application

This is a simple Todo List application built with Go. It provides a RESTful API to manage tasks.

## How to Start the App

1. Clone the repository:
    ```sh
    git clone https://github.com/manuelbeos/code-branch-todo-test.git
    cd code-branch-todo-test
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run cmd/api/main.go
    ```

4. The server will start on . http://localhost:8080

## Endpoints

### Health Check
- **GET** `/health` 
  - Response: `¡Server up!`

### Tasks 
- **POST** `/tasks` *(with random delay)*
  - Request Body:
    ```json
    {
      "title": "Task Title",
      "description": "Task Description"
    }
    ```
  - Response:
    ```json
    {
      "id": "uuid",
      "title": "Task Title",
      "description": "Task Description",
      "is_completed": false,
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

- **GET** `/tasks` *(with random delay)*
  - Response:
    ```json
    [
      {
        "id": "uuid",
        "title": "Task Title",
        "description": "Task Description",
        "is_completed": false,
        "created_at": "timestamp",
        "updated_at": "timestamp"
      }
    ]
    ```

- **GET** `/tasks/{id}`
  - Response:
    ```json
    {
      "id": "uuid",
      "title": "Task Title",
      "description": "Task Description",
      "is_completed": false,
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

- **PUT** `/tasks/{id}`
  - Request Body:
    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "is_completed": true
    }
    ```
  - Response:
    ```json
    {
      "id": "uuid",
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "is_completed": true,
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

- **DELETE** `/tasks/{id}`
  - Response: `204 No Content`


## Examples

### Create a New Task
```sh
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title": "New Task", "description": "Task Description"}'
```

### Create a New Task (PowerShell)
```sh
Invoke-WebRequest -Uri http://localhost:8080/tasks -Headers @{ "Content-Type" = "application/json" } -Method POST -Body '{"title": "New Task", "description": "Task Description"}'
```

### Get All Tasks
```sh
curl -X GET http://localhost:8080/tasks
```

### Get All Tasks (PowerShell)
```sh
Invoke-WebRequest -Uri http://localhost:8080/tasks
```

### Get Task by ID
```sh
curl -X GET http://localhost:8080/tasks/{id}
```

### Get Task by ID (PowerShell)
```sh
Invoke-WebRequest -Uri http://localhost:8080/tasks/{id}
```

### Update Task
```sh
curl -X PUT http://localhost:8080/tasks/{id} -H "Content-Type: application/json" -d '{"title": "Updated Task Title", "description": "Updated Task Description", "is_completed": true}'
```

### Update Task (PowerShell)
```sh
Invoke-WebRequest -Uri http://localhost:8080/tasks/{id} -Headers @{ "Content-Type" = "application/json" } -Method PUT -Body '{"title": "Updated Task Title", "description": "Updated Task Description", "is_completed": true}'
```

### Delete Task
```sh
curl -X DELETE http://localhost:8080/tasks/{id}
```

### Delete Task (PowerShell)
```sh
Invoke-WebRequest -Uri http://localhost:8080/tasks/{id} -Method DELETE
```