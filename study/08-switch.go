package main

import "fmt"

func main() {
	var number1 int = 200

	switch number1 {
	case 100:
		fmt.Printf("Number is 100\n")
	case 200:
		fmt.Printf("Number is 200\n")
	default:
		fmt.Printf("Number is not 100 or 200\n")
	}
}
