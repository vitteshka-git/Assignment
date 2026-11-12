package main

import "fmt"

func main() {

	name := "Vittesh"
	fmt.Println(name)

	if true {
		fmt.Println("Employee exists")
	}
}

/* 
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> go install honnef.co/go/tools/cmd/staticcheck@latest
go: downloading honnef.co/go/tools v0.8.1
go: downloading golang.org/x/tools v0.44.1-0.20260420230617-19499e7caabc
go: downloading golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678
go: downloading github.com/BurntSushi/toml v1.4.1-0.20240526193622-a339e1f7089c
go: downloading golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa
go: downloading golang.org/x/sync v0.20.0
go: downloading golang.org/x/mod v0.35.0
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> staticcheck ./...
-: # employee-managemnet/Day8_Assignment
.\2-application_logging.go:4:6: Employee redeclared in this block
        .\1-Debug.go:5:6: other declaration of Employee
.\2-application_logging.go:10:6: main redeclared in this block
        .\1-Debug.go:21:6: other declaration of main
.\3-different_log_levels.go:16:6: main redeclared in this block
        .\1-Debug.go:21:6: other declaration of main
.\4-structured_logs.go:13:6: main redeclared in this block
        .\1-Debug.go:21:6: other declaration of main
.\5-gofmt.go:5:6: main redeclared in this block
        .\1-Debug.go:21:6: other declaration of main
.\6-run_go_vet.go:5:6: main redeclared in this block
        .\1-Debug.go:21:6: other declaration of main
.\7-static_analysis.go:5:6: main redeclared in this block
        .\1-Debug.go:21:6: other declaration of main (compile)
*/