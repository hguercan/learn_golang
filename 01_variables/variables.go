package main

import "fmt"

var VisibleConstant int = 100

func main() {
	//default value will be zero
	var defaultnum int 

	//multiple variables assignment in one line
	var num1, num2 int = 12, -24 

	// variables assignment without type declaration with ':=' operator but only in loops and functions
	num3, num4 := 7, 14 

	fmt.Println("Value of defaultnum", defaultnum)
	fmt.Println("Sum of first two numbers", num1, " + ", num2, " = ", num1 + num2)
	fmt.Println("Sum of next two numbers", num3, " + ", num4, " = ", num3 + num4)
}