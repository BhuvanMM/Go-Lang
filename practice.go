package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println("Practice")

	var username string = "Bhuvan"
	fmt.Println(username)

	var age int = 20
	fmt.Println(age)

	marks := 100
	fmt.Println(marks)

	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter name")
	input, _ := read.ReadString('\n')
	fmt.Println(input)

	var fruits = []string{"Apple", "mango"}
	fruits = append(fruits, "banana")
	fmt.Println(fruits)

	marksScored := make([]int, 2)
	marksScored = append(marksScored, 100, 90)

	sort.Ints(marksScored)
	fmt.Println(marksScored)

	marksScored[0] = 50
	fmt.Println(marksScored)
}
