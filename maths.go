package main

import "fmt"
// function to add two numbers
func add(a, b int) int { 
	return a+b
}
// function to subtract two numbers
func subtract(a,b int) int {
	return a-b
}
//function to multiply two numbers
func multiply(a, b int) int {
	return a * b 
}

func main(){
	var a,b int
	// assign values 
	a = 50
	b = 5

fmt.Println("a+b=", add(a,b))
fmt.Println("a-b=", subtract(a,b))
fmt.Println("a*b=", multiply(a,b))
}
