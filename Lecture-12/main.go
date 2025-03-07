package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greet()
	add1(2, 3)

	res1 := add2(2, 3)
	fmt.Println(res1)

	res2 := add3(2, 3, 4, 5)
	fmt.Println(res2)

	bhuvan := User{"bhuvan", 20}
	fmt.Println(bhuvan)

	bhuvan.getName()

}

func greet() {
	fmt.Println("Hello Good Morning!")
}

func add1(val1 int, val2 int) {
	fmt.Println(val1 + val2)
}

func add2(val1 int, val2 int) int {
	return val1 + val2
}

func add3(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

type User struct {
	Name string
	age  int
}

func (u User) getName() {
	u.Name = "bopanna"
	fmt.Println(u.Name)
}
