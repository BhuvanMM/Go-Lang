package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user page"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter username")
	input, _ := reader.ReadString('\n')
	fmt.Println("username:", input)
}
