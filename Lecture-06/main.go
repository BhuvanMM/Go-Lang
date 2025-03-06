package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Value of ptr:", ptr)

	myNumber := 20
	var ptr = &myNumber

	fmt.Println("value of ptr:", ptr)
	fmt.Println("Value of actual ptr:", *ptr)
}
