package functions

import "fmt"

// Higher order function

var a, b int = 20, 30

// Higher order function
func processCalc (num1, num2 int, fn func(int, int)) {
	fn(num1, num2)
}

// Lower order function
func add(a, b int) {
	z := a + b
	fmt.Println("total is:::", z)
}


func LowerFn () {
	processCalc(a, b, add)
}