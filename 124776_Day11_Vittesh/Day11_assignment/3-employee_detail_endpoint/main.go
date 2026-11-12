
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
	{ID: 1, Name: "Ravi"},
	{ID: 2, Name: "Priya"},
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/employee/")
	id, err := strconv.Atoi(idText)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	for _, emp := range employees {
		if emp.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(emp)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/employee/", getEmployeeByID)
	http.ListenAndServe(":8080", nil)
}
