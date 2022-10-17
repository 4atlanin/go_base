package main

import "fmt"

func main() {
	var unexportedVariable string = "unexported variable"
	var ExportedVariable = "ExportedVariable"
	shortDeclaration := '1'
	fmt.Println(shortDeclaration)
	fmt.Println(unexportedVariable)
	fmt.Println(ExportedVariable)
}
