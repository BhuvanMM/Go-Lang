package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web services")
	performGetRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Ststus code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	//fmt.Println(string(content))
}
