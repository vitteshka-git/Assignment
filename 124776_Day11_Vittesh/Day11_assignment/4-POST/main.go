
package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var employees []Employee
var nextID = 1

func createEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var emp Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if emp.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	emp.ID = nextID
	nextID++

	employees = append(employees, emp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	http.HandleFunc("/employee", createEmployee)
	http.ListenAndServe(":8080", nil)
}
