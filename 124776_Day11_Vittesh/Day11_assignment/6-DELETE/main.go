
package main

import (
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var employees = []Employee{
	{ID: 1, Name: "Vitti"},
	{ID: 2, Name: "Priya"},
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil || id <= 0 {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/employee/", deleteEmployee)
	http.ListenAndServe(":8080", nil)
}
