
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var employees = []Employee{
	{ID: 1, Name: "vitti"},
	{ID: 2, Name: "Priya"},
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil || id <= 0 {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var updated Employee
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if updated.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	for i := range employees {
		if employees[i].ID == id {
			employees[i].Name = updated.Name

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employees[i])
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/employee/", updateEmployee)
	http.ListenAndServe(":8080", nil)
}
