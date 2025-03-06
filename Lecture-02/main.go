package main

import "fmt"

const jwtToken string = "jwtsecretkey"

func main() {
	fmt.Println("variables")

	var username string = "Bhuvan"
	fmt.Println("Username:", username)
	fmt.Printf("Variable type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type: %T\n", isLoggedIn)

	var age uint = 20
	fmt.Println(age)
	fmt.Printf("Variable type: %T\n", age)

	var name = "Bhuvan"
	fmt.Println(name)

	//name = 3 not allowed implicit declaration

	//short variable declaration
	users := 300
	fmt.Println(users)

	fmt.Println(jwtToken)
}
