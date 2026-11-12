package main

import "fmt"

func main() {

	name := "Vittesh"

	fmt.Printf("Employee: %s\n", name)
}

/* 
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> go run 6-run_go_vet.go
Employee: %!d(string=Vittesh)
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> go vet 6-run_go_vet.go
6-run_go_vet.go:9:24: fmt.Printf format %d has arg name of wrong type string
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> go vet 6-run_go_vet.go
PS C:\Users\vittesh.ka\Downloads\GO\Day8_Assignment> go run 6-run_go_vet.go
Employee: Vittesh
*/