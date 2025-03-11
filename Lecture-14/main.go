package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com"

func main() {
	fmt.Println("Handling URL")
	fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

}
