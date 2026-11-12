
package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	emp := Employee{
		ID:   1,
		Name: "Vitti",
	}

	json.NewEncoder(w).Encode(emp)
}

func main() {
	http.HandleFunc("/employee", handler)
	http.ListenAndServe(":8080", nil)
}
