
package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	http.HandleFunc("/employee", createEmployee)
	http.ListenAndServe(":8080", nil)
}
