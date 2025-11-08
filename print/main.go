package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	// Print
	// fmt.Print("Hello")
	// fmt.Print("World")

	// n, _ := fmt.Print("Hello\n")
	// fmt.Print("Written bytes: ", n, "\n")
	// fmt.Println(error)
	// Println
	// fmt.Println("Hello", "World")

	// Printf
	// fmt.Printf("Hello", "World")
	// var floatNum float64 = 12.345678

	// fmt.Printf("Formatted float: %.1f\n", floatNum)

	// Scan

	// Sprint
	name := "Julius"
	age := 28

	message := fmt.Sprintf("My name is %v and I am %d years old.", name, age)

	fmt.Println(message)

	fmt.Println(quote.Hello())

	newFunc()
}

var newVar string

func newFunc() {
	newVar = "new variable"
	fmt.Println(newVar)
}