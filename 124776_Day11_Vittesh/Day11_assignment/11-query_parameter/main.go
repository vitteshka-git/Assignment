
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "All employees"
	}

	fmt.Fprintln(w, "Search:", name)
}

func main() {
	http.HandleFunc("GET /employee", handler)
	http.ListenAndServe(":8080", nil)
}
