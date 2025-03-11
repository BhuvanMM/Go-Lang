package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lcoCourses := []course{
		{"reactjs", 299, "youtube", "abc", []string{"webdev", "js"}},
		{"angularjs", 299, "youtube", "abd", []string{"webdev", "js"}},
		{"nodejs", 299, "youtube", "abe", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, " ", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonDataFormat := []byte(
		`
		{
			"coursename": "reactjs",
			"Price": 299,
			"website": "youtube",
			"tags": [
					"webdev",
					"js"
			]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFormat)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFormat, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("json was not valid")
	}
}

func main() {
	fmt.Println("Welcome to json data")
	//EncodeJson()
	decodeJson()
}
