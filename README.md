# Task Ninja Backend

Task Ninja is a simple todo application backend built using Go (Golang), Gorilla Mux, and MongoDB.

## Features

- Create, read, update, and delete tasks
- Store tasks in a MongoDB database
- RESTful API endpoints for simple task management

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go (Golang)**: Make sure you have Go installed on your system. You can download it from the official website: [Go Downloads](https://golang.org/dl/)

- **MongoDB**: You'll need a running MongoDB server. You can install it locally or use a cloud-based service.

- **Gorilla Mux**: We'll use Gorilla Mux for routing. Install it using:

  ```
  go get -u github.com/gorilla/mux
  ```

- **MongoDB Go driver**: We'll use the official MongoDB Go driver. Install it using:

  ```
  go get go.mongodb.org/mongo-driver
  ```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/priyanshupatra02/task_ninja_backend.git
   cd task_ninja_backend
   ```

2. Set up your environment variables `.env`:

   ```

   const (
       dbHost     = "<your-host-url-copied-from-mongoDB"
       dbName     = "<as-set-in-your-mongoDB"
       collection = "<as-set-in-your-mongoDB>"
   )
   ```

3. Build and run the server:

   ```bash
   go build .
   go run main.go
   ```

## API Endpoints

- **GET /api/getAllTasks**: Get all tasks
- **POST /api/task**: Create a new task
- **PUT /api/task/{id}**: Mark an existing task as completed
- **PUT /api/undoTask/{id}**: Undo a task
- **DELETE /api/deleteTask/{id}**: Delete a task
- **DELETE /api/deleteAllTasks**: Delete all tasks

## Usage

Use a tool like `curl` or Postman to interact with the API endpoints.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to customize this README according to your project's specifics. Happy coding! 🚀
