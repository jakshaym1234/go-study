package main

import "fmt"

func main() {
	var number int = 10
	var ip *int //type pointer
	ip = &number
	fmt.Printf("Variable - %d\n", number)
	fmt.Printf("Pointer Address of Variable - %d\n", &number)
	fmt.Printf("Pointer Address - %d\n", ip) //type pointer is set for variable pointer so this is the address of the pointer
	fmt.Printf("Pointer Value - %d\n", *ip)  //this is the value of the pointer
}
