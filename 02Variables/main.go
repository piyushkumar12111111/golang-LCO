package main

import "fmt"

const loginToken = "sadfjklasdjflkjasdf" //! public variable

func main() {

	//	fmt.Println("this is variable example\n")
	var username string = "hohn doe"
	fmt.Println(username)

	fmt.Printf("Type: %T Value: %v\n", username, username)

	var isLogedin bool = false
	fmt.Println(isLogedin)

	fmt.Printf("Type: %T Value: %v\n", isLogedin, isLogedin)

	var smallInt int = 10

	fmt.Println(smallInt)

	var checking = "india"
	fmt.Printf("Type: %T Value: %v\n", checking, checking)

	valuechecking := "india"
	fmt.Printf("Type: %T Value: %v\n", valuechecking, valuechecking)

	fmt.Println(loginToken)
}
