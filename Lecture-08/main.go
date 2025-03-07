package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruits = []string{"Apple", "peach", "Tomato"}
	fmt.Printf("type of : %T\n", fruits)

	fruits = append(fruits, "Mango", "Banana")
	fmt.Println("Fruits:", fruits)

	//fruits = append(fruits[1:3])
	//fmt.Println("Fruits:", fruits)

	marks := make([]int, 4)
	marks = append(marks, 213, 100, 87, 400, 61)

	sort.Ints(marks)
	marks[0] = 10
	fmt.Println("Marks:", marks)

	var courses = []string{"reactjs", "nextjs", "swift", "nodejs", "mongodb"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
