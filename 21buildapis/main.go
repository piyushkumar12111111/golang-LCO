package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	fmt.Println("Welcome to the API")

	//! init router
	r := mux.NewRouter()

	//! mock data   //! seeding of data to the db
	cources = append(cources, Course{CourseID: "1", CourseName: "Course 1", CoursePrice: 100, Author: &Author{FullName: "Author 1", Website: "www.author1.com"}})
	cources = append(cources, Course{CourseID: "2", CourseName: "Course 2", CoursePrice: 200, Author: &Author{FullName: "Author 2", Website: "www.author2.com"}})
	cources = append(cources, Course{CourseID: "3", CourseName: "Course 3", CoursePrice: 300, Author: &Author{FullName: "Author 3", Website: "www.author3.com"}})

	//! route handlers / endpoints
	r.HandleFunc("/", ServerHome).Methods("GET")
	r.HandleFunc("/api/cources", GetAllCourses).Methods("GET")
	r.HandleFunc("/api/cources/{id}", GetOneCource).Methods("GET")
	r.HandleFunc("/api/cources", CreateOneCource).Methods("POST")
	r.HandleFunc("/api/cources/{id}", UpdateOneCource).Methods("PUT")
	r.HandleFunc("/api/cources/{id}", DeleteOneCource).Methods("DELETE")

	//! start server
	//http.ListenAndServe(":3000", r)
	log.Fatal(http.ListenAndServe(":9000", r)) //! log fatal if there is an error in the server

}

//! controller files

func ServerHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the Home Page")

	w.Write([]byte("<h1>Welcome to the API Home Page</h1>")) //! write to the response to the writer object w and convert the string to byte
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources) //! encode the courses to json and write to the response writer object w

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

	json.NewEncoder(w).Encode(course) //! successfully added the cource to the db
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

func DeleteOneCource(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	//! grab id from request
	params := mux.Vars(r)

	//! loop through the courses and find the course with the id

	for index, item := range cources {

		if item.CourseID == params["id"] {

			cources = append(cources[:index], cources[index+1:]...)

			json.NewEncoder(w).Encode("Course deleted successfully") //! crefting a response in json
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the id")

	return

}
