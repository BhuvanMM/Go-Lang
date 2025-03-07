package main

import "fmt"

func main() {
	fmt.Println("Structs")
	bhuvan := User{"Bhuvan", "bhuvan@gmail.com", true, 2000000}
	fmt.Println(bhuvan)
	fmt.Println(bhuvan.email)
}

type User struct {
	Name   string
	email  string
	status bool
	salary int
}
