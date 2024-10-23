package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

var db *sql.DB

func initDB() {
	dsn := "root:Emotib@t15.@tcp(mariadb:3306)/API_REST"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}

	// Verifica si la conexión es exitosa
	if err := db.Ping(); err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
}

var tasks []task

func getTasks(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(tasks)

	if db == nil {
		http.Error(w, "Database connection is not initialized", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	tasks := []task{}

	rows, err := db.Query("SELECT id_act, nombre, contenido FROM actividades")
	if err != nil {
		log.Println("Error querying database:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t task
		if err := rows.Scan(&t.ID, &t.Name, &t.Content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Elimina la tarea de la base de datos
	_, err = db.Exec("DELETE FROM actividades WHERE id_act = ?", taskID)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // No content

}

func createTasks(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &newTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Inserta en la base de datos
	result, err := db.Exec("INSERT INTO actividades (nombre, contenido) VALUES (?, ?)", newTask.Name, newTask.Content)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	lastID, err := result.LastInsertId() // Obtener el ID del nuevo registro
	if err != nil {
		http.Error(w, "Error retrieving last insert ID", http.StatusInternalServerError)
		return
	}
	newTask.ID = int(lastID) // Conversión de int64 a int

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var t task
	err = db.QueryRow("SELECT id_act, nombre, contenido FROM actividades WHERE id_act = ?", taskID).Scan(&t.ID, &t.Name, &t.Content)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &updatedTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Actualiza la tarea en la base de datos
	_, err = db.Exec("UPDATE actividades SET nombre = ?, contenido = ? WHERE id_act = ?", updatedTask.Name, updatedTask.Content, taskID)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	updatedTask.ID = taskID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Api Chest")
}

func main() {
	initDB()
	defer db.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/newtask", createTasks).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
