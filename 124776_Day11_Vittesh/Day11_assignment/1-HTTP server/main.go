
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Employee API")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on http://localhost:10005")
	http.ListenAndServe(":10005", nil)
}
