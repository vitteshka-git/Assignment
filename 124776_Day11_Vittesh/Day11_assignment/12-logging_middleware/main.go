
package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Println(
			"Method:", r.Method,
			"Path:", r.URL.Path,
			"Duration:", time.Since(start),
		)
	})
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Employee API"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee", employeeHandler)

	server := loggingMiddleware(mux)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
