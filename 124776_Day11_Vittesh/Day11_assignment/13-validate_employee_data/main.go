
package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Employee struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee

	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(emp.Name) == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if !strings.Contains(emp.Email, "@") {
		http.Error(w, "Valid email is required", http.StatusBadRequest)
		return
	}

	if emp.Salary <= 0 {
		http.Error(w, "Salary must be greater than zero", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	http.HandleFunc("POST /employee", createEmployee)
	http.ListenAndServe(":8080", nil)
}
