package closure

import "fmt"

var age = 20

func outerFn (age int) func(int) {

	var money int = 50

	innerFn := func(incNum int) {
		if age >= 18 {
			fmt.Println("Current money::: ", money + incNum)
		} else if money < incNum {
			fmt.Println("Current money::: ", money + incNum)
		} else {
			fmt.Println("Current money::: ", money + incNum)
		}
	}
	return innerFn
}

func FirstFn() {
	callInner := outerFn(age)
	
	callInner(20)
	callInner(30)

	callInner2 := outerFn(age)
	callInner2(20)
	callInner2(30)
}