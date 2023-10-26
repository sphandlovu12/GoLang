// Restful api for managing list of tasks
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Task struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tasks []Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, task := range tasks {
		if task.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Task not found")
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = fmt.Sprint(len(tasks) + 1)
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func main() {
	router := mux.NewRouter()

	// Sample data for testing
	tasks = append(tasks, Task{ID: "1", Title: "Task 1", Content: "This is the content of Task 1"})
	tasks = append(tasks, Task{ID: "2", Title: "Task 2", Content: "This is the content of Task 2"})

	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
