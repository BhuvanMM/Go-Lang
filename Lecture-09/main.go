package main

import "fmt"

func main() {
	langs := make(map[string]string)

	langs["js"] = "javascript"
	langs["rb"] = "ruby"
	langs["py"] = "python"

	fmt.Println(langs)
	fmt.Println(langs["js"])

	delete(langs, "rb")
	fmt.Println(langs)

	for key, value := range langs {
		fmt.Printf("Key is %v and value is %v\n", key, value)
	}
}
