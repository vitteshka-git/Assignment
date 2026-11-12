
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

type App struct {
	db *sql.DB
}

// Send JSON response.
func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Logging middleware.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		log.Println(
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

// Validate employee data.
func validateEmployee(emp Employee) string {
	if strings.TrimSpace(emp.Name) == "" {
		return "Name is required"
	}

	if !strings.Contains(emp.Email, "@") {
		return "Valid email is required"
	}

	if emp.Salary <= 0 {
		return "Salary must be greater than zero"
	}

	return ""
}

// GET /employees
// GET /employees?name=Ravi
func (a *App) getEmployees(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.URL.Query().Get("name"))

	query := `SELECT id, name, email, salary
	          FROM employees`

	var rows *sql.Rows
	var err error

	if name != "" {
		query += ` WHERE name ILIKE $1 ORDER BY id`
		rows, err = a.db.Query(query, "%"+name+"%")
	} else {
		query += ` ORDER BY id`
		rows, err = a.db.Query(query)
	}

	if err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not retrieve employees"})
		return
	}
	defer rows.Close()

	employees := []Employee{}

	for rows.Next() {
		var emp Employee

		if err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Salary,
		); err != nil {
			sendJSON(w, http.StatusInternalServerError,
				map[string]string{"error": "Could not read employees"})
			return
		}

		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not read employees"})
		return
	}

	sendJSON(w, http.StatusOK, employees)
}

// GET /employees/{id}
func (a *App) getEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		sendJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Invalid employee ID"})
		return
	}

	var emp Employee

	err = a.db.QueryRow(
		`SELECT id, name, email, salary
		 FROM employees WHERE id = $1`,
		id,
	).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Salary)

	if err == sql.ErrNoRows {
		sendJSON(w, http.StatusNotFound,
			map[string]string{"error": "Employee not found"})
		return
	}

	if err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not retrieve employee"})
		return
	}

	sendJSON(w, http.StatusOK, emp)
}

// POST /employees
func (a *App) createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee

	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Invalid JSON"})
		return
	}

	if message := validateEmployee(emp); message != "" {
		sendJSON(w, http.StatusBadRequest,
			map[string]string{"error": message})
		return
	}

	err := a.db.QueryRow(
		`INSERT INTO employees (name, email, salary)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		strings.TrimSpace(emp.Name),
		strings.TrimSpace(emp.Email),
		emp.Salary,
	).Scan(&emp.ID)

	if err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not create employee"})
		return
	}

	sendJSON(w, http.StatusCreated, emp)
}

// PUT /employees/{id}
func (a *App) updateEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		sendJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Invalid employee ID"})
		return
	}

	var emp Employee

	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		sendJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Invalid JSON"})
		return
	}

	if message := validateEmployee(emp); message != "" {
		sendJSON(w, http.StatusBadRequest,
			map[string]string{"error": message})
		return
	}

	result, err := a.db.Exec(
		`UPDATE employees
		 SET name = $1, email = $2, salary = $3
		 WHERE id = $4`,
		strings.TrimSpace(emp.Name),
		strings.TrimSpace(emp.Email),
		emp.Salary,
		id,
	)

	if err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not update employee"})
		return
	}

	count, err := result.RowsAffected()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not verify update"})
		return
	}

	if count == 0 {
		sendJSON(w, http.StatusNotFound,
			map[string]string{"error": "Employee not found"})
		return
	}

	emp.ID = id
	sendJSON(w, http.StatusOK, emp)
}

// DELETE /employees/{id}
func (a *App) deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		sendJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Invalid employee ID"})
		return
	}

	result, err := a.db.Exec(
		`DELETE FROM employees WHERE id = $1`,
		id,
	)

	if err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not delete employee"})
		return
	}

	count, err := result.RowsAffected()
	if err != nil {
		sendJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not verify deletion"})
		return
	}

	if count == 0 {
		sendJSON(w, http.StatusNotFound,
			map[string]string{"error": "Employee not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("Set the DB_PASSWORD environment variable")
	}

	connStr := "host=localhost port=5432 user=postgres password=Root dbname=employee1 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database configuration failed:", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("PostgreSQL connected")

	app := &App{db: db}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /employees", app.getEmployees)
	mux.HandleFunc("GET /employees/{id}", app.getEmployee)
	mux.HandleFunc("POST /employees", app.createEmployee)
	mux.HandleFunc("PUT /employees/{id}", app.updateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", app.deleteEmployee)

	server := loggingMiddleware(mux)

	log.Println("Employee API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
