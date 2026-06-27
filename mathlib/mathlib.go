package mathlib


func Calculation(num1, num2 int, logic string) int {
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