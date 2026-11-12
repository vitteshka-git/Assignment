# Employee Microservices Architecture

This codebase implements all 15 tasks for Day 18 with comments placed directly above each implementation logic block.

## How to Run:
1. Ensure PostgreSQL is running and databases are created (`employee_db`, `department_db`, `auth_db`).
2. Run each service in separate terminals:
   - `go run ./services/department`
   - `go run ./services/employee`
   - `go run ./services/auth`
   - `go run ./gateway`
