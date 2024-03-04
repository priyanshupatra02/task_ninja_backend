package router

import (
	"github.com/gorilla/mux"
	"github.com/priyanshupatra02/task-ninja-backend/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// get all tasks
	router.HandleFunc("/api/getAllTasks", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	//create a task
	router.HandleFunc("/api/task", middleware.CreateATask).Methods("POST", "OPTIONS")
	//update a task
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	//undo a task
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	//delete a task
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteATask).Methods("DELETE", "OPTIONS")
	//delete all tasks
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")

	return router
}
