package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("Handling HTTP Requests")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	//fmt.Println(response)

	databody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(databody))

	response.Body.Close()
}
