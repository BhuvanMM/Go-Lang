package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware - file
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

// controllers - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Home Page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, value := range courses {
		if value.CourseID == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found for the given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a new course")
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses := append(courses, course)
}

func main() {
	fmt.Println("Building API")
}
