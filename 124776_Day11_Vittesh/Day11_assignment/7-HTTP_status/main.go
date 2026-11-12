
package main

import (
	"encoding/json"
	"net/http"
)

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Employees retrieved successfully",
		})

	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Employee created successfully",
		})

	case http.MethodDelete:
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method not allowed",
		})
	}
}

func main() {
	http.HandleFunc("/employee", employeeHandler)
	http.ListenAndServe(":8080", nil)
}
