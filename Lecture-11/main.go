package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Conditional statements")

	loginCount := 10
	var result string

	if loginCount < 5 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Dangerous User"
	} else {
		result = "Good User"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Number is small")
	}

	fmt.Println("Switch Cases")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	default:
		fmt.Println("Default")
	}

	fmt.Println("Loops in Go")

	days := []string{"sunday", "monday", "tuesday", "saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, days := range days {
		fmt.Printf("%v and %v\n", index, days)
	}

	idx := 1
	for idx < 10 {
		fmt.Println(idx)
		idx++
	}

}
