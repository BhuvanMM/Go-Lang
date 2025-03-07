package main

import "fmt"

func main() {
	fmt.Println("Arrays in go:")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[2] = "Banana"
	fruits[3] = "Cherry"

	fmt.Println("Fruits array is:", fruits)
	fmt.Println("Length:", len(fruits))

	var veggies = [3]string{"Tomato", "beans", "mushroom"}
	fmt.Println("Veggies:", veggies)
}
