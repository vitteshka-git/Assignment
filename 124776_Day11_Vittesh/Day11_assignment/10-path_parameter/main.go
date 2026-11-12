
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")

	id, err := strconv.Atoi(idText)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Employee ID: %d", id)
}

func main() {
	http.HandleFunc("GET /employee/{id}", handler)
	http.ListenAndServe(":8080", nil)
}
