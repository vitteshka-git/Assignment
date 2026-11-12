
package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	employees := []Employee{
		{ID: 1, Name: "Ravi", Email: "ravi@example.com", Salary: 30000},
		{ID: 2, Name: "Priya", Email: "priya@example.com", Salary: 40000},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	http.HandleFunc("/employee", getEmployees)
	http.ListenAndServe(":8080", nil)
}
