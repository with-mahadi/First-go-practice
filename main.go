package main

import (
	"fmt"

	"example.com/functions"
)

/*
	1. For custom package use another folder should run this command
	2. go mod init example.com
*/
func main() {
	// num1, num2, logic := getUserInput()
	// value := mathlib.Calculation(num1, num2, logic)
	// user_name := getUserName()
	
	// fmt.Println(user_name, "Your calculation result is::", value)

	fmt.Println("Execute the higher order function")
	functions.LowerFn()
}

func getUserInput() (int, int, string) {
	var num1, num2 int
	var logic string
	fmt.Println("Enter first number")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number")
	fmt.Scanln(&num2)

	fmt.Println("What do you want, (e.g. sum,sub,mul,div)")
	fmt.Scanln(&logic)

	return num1, num2, logic
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name here::")
	fmt.Scanln(&name)
	return name
}
