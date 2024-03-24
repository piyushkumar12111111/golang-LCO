package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Hello mod in golang")

	greeter()

	r := mux.NewRouter()
	r.HandleFunc(
		"/", ServerHome,
	).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))

}

func greeter() {

	fmt.Println("Namstey from golang")
}

func ServerHome(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("Server home")

	w.Write([]byte("<h1>Hello from Server Home</h1>"))
}
