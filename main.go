package main

import "fmt"

func main() {
	num1, num2, logic := getUserInput()
	value := calculation(num1, num2, logic)
	user_name := getUserName()
	
	fmt.Println(user_name, "Your calculation result is::", value)
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
func calculation(num1, num2 int, logic string) int {
	switch logic {
		case "sum":
			return  num1 + num2
		case "sub":
			return num1 - num2
		case "mul":
			return num1 * num2
		case "div": 
			return num1 / num2
		default: 
			return num1 + num2		
	}
}