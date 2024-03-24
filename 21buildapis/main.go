package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//! model for course

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

//! fake db

var cources []Course

//! middleware, helper - file

func (c *Course) Isempty() bool {

	return c.CourseID == "" && c.CourseName == ""
}

func main() {

}

//! controller files

func ServerHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the Home Page")

	w.Write([]byte("<h1>Welcome to the API Home Page</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)

}

func GetOneCource(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//! grab id from request
	params := mux.Vars(r)

	//! loop through the courses and find the course with the id
	for _, item := range cources {

		if item.CourseID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the id")
	return

}

func CreateOneCource(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//! if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//! if {course_id, course_name} is empty
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.Isempty() {
		json.NewEncoder(w).Encode("Course is empty")
		return
	}
	//	rand.Seed(time.Now().UnixNano())
	//	course.CourseID = strconv.Itoa(rand.Intn(1000))
	cources = append(cources, course)

	json.NewEncoder(w).Encode(course)
	return
}

func UpdateOneCource(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//! grab id from request
	params := mux.Vars(r)

	//! loop through the courses and find the course with the id
	for index, item := range cources {

		if item.CourseID == params["id"] {

			cources = append(cources[:index], cources[index+1:]...)

			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseID = params["id"]

			cources = append(cources, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the id")
	return
}
