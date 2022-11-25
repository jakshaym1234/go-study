package main

import "fmt"

func main() {
	var num1 int = 100
	var num2 int = 200
	var returnvalue int
	returnvalue = getMax(num1, num2)
	fmt.Printf("%d\n", returnvalue)
}

func getMax(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
