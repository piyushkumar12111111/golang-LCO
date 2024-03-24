package main

import "fmt"

func main() {

	piyush := User{
		Name:  "Piyush",
		Email: "piyushkumar@go.dev",
		Age:   25,
	}
	fmt.Println(piyush)
	fmt.Printf("User Name is %v and Email is %v and Age is %v\n", piyush.Name, piyush.Email, piyush.Age)
}

type User struct {
	Name  string
	Email string
	Age   int
}

//! There is no concept of classes in Go there is only struct and methods and no inheritance
